package githubridgeclient

import (
	"context"

	pb "github.com/brotherlogic/githubridge/proto"
)

type TestClient struct {
	counter int64
	issues  map[string]int64
}

func GetTestClient() GithubridgeClient {
	return &TestClient{issues: make(map[string]int64), counter: 0}
}

func (c *TestClient) CreateIssue(ctx context.Context, req *pb.CreateIssueRequest) (*pb.CreateIssueResponse, error) {
	c.counter++
	c.issues[req.GetTitle()] = c.counter
	return &pb.CreateIssueResponse{IssueId: c.counter}, nil
}

func (c *TestClient) CloseIssue(ctx context.Context, req *pb.CloseIssueRequest) (*pb.CloseIssueResponse, error) {
	return &pb.CloseIssueResponse{}, nil
}

func (c *TestClient) CommentOnIssue(ctx context.Context, req *pb.CommentOnIssueRequest) (*pb.CommentOnIssueResponse, error) {
	return &pb.CommentOnIssueResponse{}, nil
}

func (c *TestClient) GetIssue(ctx context.Context, req *pb.GetIssueRequest) (*pb.GetIssueResponse, error) {
	return &pb.GetIssueResponse{}, nil
}
