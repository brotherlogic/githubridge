package githubridgeclient

import (
	"context"
	"fmt"

	pb "github.com/brotherlogic/githubridge/proto"
)

type TestClient struct {
	counter int64
	issues  map[int64]string
	labels  map[string][]string
}

func GetTestClient() GithubridgeClient {
	return &TestClient{
		issues:  make(map[int64]string),
		counter: 0,
		labels:  make(map[string][]string),
	}
}

func (c *TestClient) AddLabel(ctx context.Context, req *pb.AddLabelRequest) (*pb.AddLabelResponse, error) {
	label := fmt.Sprintf("%v-%v-%v", req.GetUser(), req.GetRepo(), req.GetId())
	c.labels[label] = append(c.labels[label], req.GetLabel())
	return &pb.AddLabelResponse{}, nil
}

func (c *TestClient) CreateIssue(ctx context.Context, req *pb.CreateIssueRequest) (*pb.CreateIssueResponse, error) {
	c.counter++
	c.issues[c.counter] = req.GetTitle()
	return &pb.CreateIssueResponse{IssueId: c.counter}, nil
}

func (c *TestClient) CloseIssue(ctx context.Context, req *pb.CloseIssueRequest) (*pb.CloseIssueResponse, error) {
	delete(c.issues, req.GetId())
	return &pb.CloseIssueResponse{}, nil
}

func (c *TestClient) CommentOnIssue(ctx context.Context, req *pb.CommentOnIssueRequest) (*pb.CommentOnIssueResponse, error) {
	return &pb.CommentOnIssueResponse{}, nil
}

func (c *TestClient) GetIssue(ctx context.Context, req *pb.GetIssueRequest) (*pb.GetIssueResponse, error) {
	return &pb.GetIssueResponse{}, nil
}

func (c *TestClient) GetLabels(ctx context.Context, req *pb.GetLabelsRequest) (*pb.GetLabelsResponse, error) {
	return &pb.GetLabelsResponse{Labels: c.labels[fmt.Sprintf("%v-%v-%v", req.GetUser(), req.GetRepo(), req.GetId())]}, nil
}

func (c *TestClient) GetIssues(ctx context.Context, req *pb.GetIssuesRequest) (*pb.GetIssuesResponse, error) {
	var issues []*pb.GithubIssue
	for c, issue := range c.issues {
		issues = append(issues, &pb.GithubIssue{Id: int64(c), Title: issue})
	}
	return &pb.GetIssuesResponse{Issues: issues}, nil
}
