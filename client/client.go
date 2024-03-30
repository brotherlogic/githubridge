package githubridgeclient

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	pb "github.com/brotherlogic/githubridge/proto"
)

type GithubridgeClient interface {
	CreateIssue(ctx context.Context, req *pb.CreateIssueRequest) (*pb.CreateIssueResponse, error)
	CloseIssue(ctx context.Context, req *pb.CloseIssueRequest) (*pb.CloseIssueResponse, error)
	CommentOnIssue(ctx context.Context, req *pb.CommentOnIssueRequest) (*pb.CommentOnIssueResponse, error)
	GetIssue(ctx context.Context, req *pb.GetIssueRequest) (*pb.GetIssueResponse, error)
	GetLabels(ctx context.Context, req *pb.GetLabelsRequest) (*pb.GetLabelsResponse, error)
	GetIssues(ctx context.Context, req *pb.GetIssuesRequest) (*pb.GetIssuesResponse, error)
}

type rClient struct {
	gClient pb.GithubridgeServiceClient
	passkey string
}

func GetClientInternal(pass string) (GithubridgeClient, error) {
	return getClientInternal(pass, "githubridge.githubridge:8080")
}

func GetClientExternal(pass string) (GithubridgeClient, error) {
	return getClientInternal(pass, "ghb.brotherlogic-backend.com:80")
}

func getClientInternal(pass string, address string) (GithubridgeClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &rClient{gClient: pb.NewGithubridgeServiceClient(conn), passkey: pass}, nil
}

func (c *rClient) CreateIssue(ctx context.Context, req *pb.CreateIssueRequest) (*pb.CreateIssueResponse, error) {
	nctx := metadata.AppendToOutgoingContext(context.Background(), "auth-token", string(c.passkey))
	return c.gClient.CreateIssue(nctx, req)
}

func (c *rClient) CloseIssue(ctx context.Context, req *pb.CloseIssueRequest) (*pb.CloseIssueResponse, error) {
	nctx := metadata.AppendToOutgoingContext(context.Background(), "auth-token", string(c.passkey))
	return c.gClient.CloseIssue(nctx, req)
}
func (c *rClient) CommentOnIssue(ctx context.Context, req *pb.CommentOnIssueRequest) (*pb.CommentOnIssueResponse, error) {
	nctx := metadata.AppendToOutgoingContext(context.Background(), "auth-token", string(c.passkey))
	return c.gClient.CommentOnIssue(nctx, req)
}
func (c *rClient) GetIssue(ctx context.Context, req *pb.GetIssueRequest) (*pb.GetIssueResponse, error) {
	nctx := metadata.AppendToOutgoingContext(context.Background(), "auth-token", string(c.passkey))
	return c.gClient.GetIssue(nctx, req)
}
func (c *rClient) GetLabels(ctx context.Context, req *pb.GetLabelsRequest) (*pb.GetLabelsResponse, error) {
	nctx := metadata.AppendToOutgoingContext(context.Background(), "auth-token", string(c.passkey))
	return c.gClient.GetLabels(nctx, req)
}
func (c *rClient) GetIssues(ctx context.Context, req *pb.GetIssuesRequest) (*pb.GetIssuesResponse, error) {
	nctx := metadata.AppendToOutgoingContext(context.Background(), "auth-token", string(c.passkey))
	return c.gClient.GetIssues(nctx, req)
}
