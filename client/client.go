package githubridgeclient

import (
	"context"

	"google.golang.org/grpc"

	pb "github.com/brotherlogic/githubridge/proto"
)

type GithubridgeClient interface {
	AddIssue(ctx context.Context, req *pb.AddIssueRequest) (*pb.AddIssueResponse, error)
}

type rClient struct {
	gClient pb.GithubBridgeClient
}

func GetClient() (GithubridgeClient, error) {
	conn, err := grpc.Dial("githubridge.githubridge:8080", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &rClient{gClient: pb.NewGithubBridgeClient(conn)}, nil
}

func (c *rClient) AddIssue(ctx context.Context, req *pb.AddIssueRequest) (*pb.AddIssueResponse, error) {
	return c.gClient.AddIssue(ctx, req)
}
