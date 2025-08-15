package server

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v74/github"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/brotherlogic/githubridge/proto"
)

var (
	creates = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "githubridge_creates",
		Help: "The number of repos being tracked",
	}, []string{"repo", "code"})

	serverRequests = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "githubridge_requests",
		Help: "The number of server requests",
	}, []string{"method", "status"})
)

func (s *Server) ServerInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	h, err := handler(ctx, req)
	serverRequests.With(prometheus.Labels{"status": status.Convert(err).Code().String(), "method": info.FullMethod}).Inc()
	return h, err
}

func (s *Server) CreateIssue(ctx context.Context, req *pb.CreateIssueRequest) (*pb.CreateIssueResponse, error) {
	defer s.metrics()

	// Fail if an issue is open with that name
	log.Printf("Checking %v issues for dedup", len(s.issues))
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
	processResponse(resp, "issues-create")
	if err != nil {
		creates.With(prometheus.Labels{"repo": req.GetRepo(), "code": fmt.Sprintf("%v", status.Code(err))}).Inc()
		return nil, err
	}

	creates.With(prometheus.Labels{"repo": req.GetRepo(), "code": fmt.Sprintf("%v", resp.StatusCode)}).Inc()
	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("Bad response code: %v", resp.StatusCode)
	}

	// Add this issue to the exisitng list
	s.issues = append(s.issues, &pb.GithubIssue{
		Repo:  req.GetRepo(),
		User:  req.GetUser(),
		Id:    int64(issue.GetNumber()),
		Title: req.GetTitle(),
		State: pb.IssueState_ISSUE_STATE_OPEN,
	})

	return &pb.CreateIssueResponse{IssueId: (int64(issue.GetNumber()))}, nil
}

func (s *Server) CloseIssue(ctx context.Context, req *pb.CloseIssueRequest) (*pb.CloseIssueResponse, error) {

	_, resp, err := s.client.Issues.Edit(ctx, req.GetUser(), req.GetRepo(), int(req.GetId()), &github.IssueRequest{
		State: proto.String("closed"),
	})
	processResponse(resp, "issues-edit")

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
	processResponse(resp, "issues-get")

	if resp != nil && resp.StatusCode == 404 {
		return nil, status.Errorf(codes.NotFound, "Cannot find %v/%v/%v", req.User, req.GetRepo(), req.GetId())
	}

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Bad response code: %v", resp.StatusCode)
	}

	var labels []string
	for _, label := range issue.Labels {
		labels = append(labels, label.GetName())
	}

	subissues, resp, err := s.client.SubIssue.ListByIssue(ctx, req.GetUser(), req.GetRepo(), int64(req.GetId()), &github.IssueListOptions{})
	processResponse(resp, "subissues-listbyissue")

	var subs []*pb.GithubIssue
	for _, subissue := range subissues {
		subs = append(subs, &pb.GithubIssue{
			Id:   *subissue.ID,
			Repo: *subissue.Repository.Name,
			User: *subissue.User.Name,
		})
	}

	return &pb.GetIssueResponse{
		State:     convertIssueState(issue.GetState()),
		Comments:  int32(issue.GetComments()),
		Labels:    labels,
		SubIssues: subs,
	}, nil
}

func (s *Server) GetIssues(ctx context.Context, req *pb.GetIssuesRequest) (*pb.GetIssuesResponse, error) {
	return &pb.GetIssuesResponse{Issues: s.issues}, nil
}
