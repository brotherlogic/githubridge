package githubridgeclient

import (
	"context"

	pb "github.com/brotherlogic/githubridge/proto"
)

type TestClient struct {
	counter int32
	issues  map[string]int32
}

func GetTestClient() GithubridgeClient {
	return &TestClient{issues: make(map[string]int32), counter: 0}
}

func (c *TestClient) AddIssue(ctx context.Context, req *pb.AddIssueRequest) (*pb.AddIssueResponse, error) {
	c.counter++
	c.issues[req.GetTitle()] = c.counter
	return &pb.AddIssueResponse{IssueId: c.counter}, nil
}
