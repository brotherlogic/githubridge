package server

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v50/github"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/brotherlogic/githubridge/proto"
)

type Server struct {
	client *github.Client
	repos  []string
	user   string
	issues []*pb.GithubIssue
	ready  bool // ready to server
}

func NewServer(client *github.Client, user string) *Server {
	s := &Server{client: client, user: user, ready: true}
	err := s.startup(context.Background())
	if err != nil {
		log.Printf("Failed startup: %v", err)
		os.Exit(1)
	}
	s.ready = true
	return s
}

func (s *Server) CreateIssue(ctx context.Context, req *pb.CreateIssueRequest) (*pb.CreateIssueResponse, error) {
	issue, resp, err := s.client.Issues.Create(ctx, req.GetUser(), req.GetRepo(), &github.IssueRequest{
		Title: proto.String(req.GetTitle()),
		Body:  proto.String(req.GetBody()),
	})
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("Bad response code: %v", resp.StatusCode)
	}

	return &pb.CreateIssueResponse{IssueId: (int64(issue.GetNumber()))}, nil
}

func (s *Server) CloseIssue(ctx context.Context, req *pb.CloseIssueRequest) (*pb.CloseIssueResponse, error) {

	_, resp, err := s.client.Issues.Edit(ctx, req.GetUser(), req.GetRepo(), int(req.GetId()), &github.IssueRequest{
		State: proto.String("closed"),
	})

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Bad response code: %v", resp.StatusCode)
	}

	return &pb.CloseIssueResponse{}, nil
}

func (s *Server) CommentOnIssue(ctx context.Context, req *pb.CommentOnIssueRequest) (*pb.CommentOnIssueResponse, error) {
	_, resp, err := s.client.Issues.CreateComment(ctx, req.GetUser(), req.GetRepo(), int(req.GetId()), &github.IssueComment{
		Body: proto.String(req.GetComment()),
	})

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("bad response code for comment: %v", resp.StatusCode)
	}

	return &pb.CommentOnIssueResponse{}, nil
}

func (s *Server) GetIssue(ctx context.Context, req *pb.GetIssueRequest) (*pb.GetIssueResponse, error) {
	issue, resp, err := s.client.Issues.Get(ctx, req.GetUser(), req.GetRepo(), int(req.GetId()))

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
