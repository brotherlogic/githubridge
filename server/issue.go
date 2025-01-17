package server

import (
	"context"
	"fmt"
	"log"

	pb "github.com/brotherlogic/githubridge/proto"

	"github.com/google/go-github/github"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func (s *Server) CreateIssue(ctx context.Context, req *pb.CreateIssueRequest) (*pb.CreateIssueResponse, error) {
	// Fail if an issue is open with that name
	for _, issue := range s.issues {
		if issue.GetTitle() == req.GetTitle() && issue.GetState() == pb.IssueState_ISSUE_STATE_OPEN {
			log.Printf("Returning error with issue: %v", issue)
			return &pb.CreateIssueResponse{IssueId: issue.GetId(), AlreadyExistingIssue: true}, nil
		}
	}

	issue, resp, err := s.client.Issues.Create(ctx, req.GetUser(), req.GetRepo(), &github.IssueRequest{
		Title: proto.String(req.GetTitle()),
		Body:  proto.String(req.GetBody()),
	})
	processResponse(resp)
	if err != nil {
		creates.With(prometheus.Labels{"repo": req.GetRepo(), "code": fmt.Sprintf("%v", status.Code(err))}).Inc()
		return nil, err
	}

	creates.With(prometheus.Labels{"repo": req.GetRepo(), "code": fmt.Sprintf("%v", resp.StatusCode)}).Inc()
	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("Bad response code: %v", resp.StatusCode)
	}

	return &pb.CreateIssueResponse{IssueId: (int64(issue.GetNumber()))}, nil
}

func (s *Server) CloseIssue(ctx context.Context, req *pb.CloseIssueRequest) (*pb.CloseIssueResponse, error) {

	_, resp, err := s.client.Issues.Edit(ctx, req.GetUser(), req.GetRepo(), int(req.GetId()), &github.IssueRequest{
		State: proto.String("closed"),
	})
	processResponse(resp)

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Bad response code: %v", resp.StatusCode)
	}

	return &pb.CloseIssueResponse{}, nil
}

func (s *Server) GetIssue(ctx context.Context, req *pb.GetIssueRequest) (*pb.GetIssueResponse, error) {
	issue, resp, err := s.client.Issues.Get(ctx, req.GetUser(), req.GetRepo(), int(req.GetId()))
	processResponse(resp)

	if resp != nil && resp.StatusCode == 404 {
		return nil, status.Errorf(codes.NotFound, "Cannot find %v/%v/%v", req.User, req.GetRepo(), req.GetId())
	}

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Bad response code: %v", resp.StatusCode)
	}
	return &pb.GetIssueResponse{
		State:    issue.GetState(),
		Comments: int32(issue.GetComments()),
	}, nil
}

func (s *Server) GetIssues(ctx context.Context, req *pb.GetIssuesRequest) (*pb.GetIssuesResponse, error) {
	return &pb.GetIssuesResponse{Issues: s.issues}, nil
}
