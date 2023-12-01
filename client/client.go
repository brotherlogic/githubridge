package githubridgeclient

import (
	"context"

	"google.golang.org/grpc"

	pb "github.com/brotherlogic/githubridge/proto"
)

type GithubridgeClient interface {
	CreateIssue(ctx context.Context, req *pb.CreateIssueRequest) (*pb.CreateIssueResponse, error)
	CloseIssue(ctx context.Context, req *pb.CloseIssueRequest) (*pb.CloseIssueResponse, error)
	CommentOnIssue(ctx context.Context, req *pb.CommentOnIssueRequest) (*pb.CommentOnIssueResponse, error)
	GetIssue(ctx context.Context, req *pb.GetIssueRequest) (*pb.GetIssueResponse, error)
}

type rClient struct {
	gClient pb.GithubridgeServiceClient
}

func GetClient() (GithubridgeClient, error) {
	conn, err := grpc.Dial("githubridge.githubridge:8080", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &rClient{gClient: pb.NewGithubridgeServiceClient(conn)}, nil
}

func (c *rClient) CreateIssue(ctx context.Context, req *pb.CreateIssueRequest) (*pb.CreateIssueResponse, error) {
	return c.gClient.CreateIssue(ctx, req)
}

func (c *rClient) CloseIssue(ctx context.Context, req *pb.CloseIssueRequest) (*pb.CloseIssueResponse, error) {
	return c.gClient.CloseIssue(ctx, req)
}
func (c *rClient) CommentOnIssue(ctx context.Context, req *pb.CommentOnIssueRequest) (*pb.CommentOnIssueResponse, error) {
	return c.gClient.CommentOnIssue(ctx, req)
}
func (c *rClient) GetIssue(ctx context.Context, req *pb.GetIssueRequest) (*pb.GetIssueResponse, error) {
	return c.gClient.GetIssue(ctx, req)
}
