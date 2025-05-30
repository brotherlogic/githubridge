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
	AddLabel(ctx context.Context, req *pb.AddLabelRequest) (*pb.AddLabelResponse, error)
	DeleteLabel(ctx context.Context, req *pb.DeleteLabelRequest) (*pb.DeleteLabelResponse, error)
	GetComments(ctx context.Context, req *pb.GetCommentsRequest) (*pb.GetCommentsResponse, error)
	GetRepos(ctx context.Context, req *pb.GetReposRequest) (*pb.GetReposResponse, error)
	GetRepo(ctx context.Context, req *pb.GetRepoRequest) (*pb.GetRepoResponse, error)
	ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error)
	GetFile(ctx context.Context, req *pb.GetFileRequest) (*pb.GetFileResponse, error)
	UpdateFile(ctx context.Context, req *pb.UpdateFileRequest) (*pb.UpdateFileResponse, error)
	GetProjects(ctx context.Context, req *pb.GetProjectsRequest) (*pb.GetProjectsResponse, error)
	GetReleases(ctx context.Context, req *pb.GetReleasesRequest) (*pb.GetReleasesResponse, error)
}

type rClient struct {
	gClient pb.GithubridgeServiceClient
	passkey string
}

func GetClientInternal() (GithubridgeClient, error) {
	return getClientInternal("", "githubridge.githubridge:8082")
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

func (c *rClient) AddLabel(ctx context.Context, req *pb.AddLabelRequest) (*pb.AddLabelResponse, error) {
	nctx := metadata.AppendToOutgoingContext(context.Background(), "auth-token", string(c.passkey))
	return c.gClient.AddLabel(nctx, req)
}

func (c *rClient) CreateIssue(ctx context.Context, req *pb.CreateIssueRequest) (*pb.CreateIssueResponse, error) {
	nctx := metadata.AppendToOutgoingContext(context.Background(), "auth-token", string(c.passkey))
	return c.gClient.CreateIssue(nctx, req)
}

func (c *rClient) GetComments(ctx context.Context, req *pb.GetCommentsRequest) (*pb.GetCommentsResponse, error) {
	nctx := metadata.AppendToOutgoingContext(context.Background(), "auth-token", string(c.passkey))
	return c.gClient.GetComments(nctx, req)
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
func (c *rClient) DeleteLabel(ctx context.Context, req *pb.DeleteLabelRequest) (*pb.DeleteLabelResponse, error) {
	nctx := metadata.AppendToOutgoingContext(context.Background(), "auth-token", string(c.passkey))
	return c.gClient.DeleteLabel(nctx, req)
}

func (c *rClient) GetRepos(ctx context.Context, req *pb.GetReposRequest) (*pb.GetReposResponse, error) {
	nctx := metadata.AppendToOutgoingContext(context.Background(), "auth-token", string(c.passkey))
	return c.gClient.GetRepos(nctx, req)
}
func (c *rClient) GetRepo(ctx context.Context, req *pb.GetRepoRequest) (*pb.GetRepoResponse, error) {
	nctx := metadata.AppendToOutgoingContext(context.Background(), "auth-token", string(c.passkey))
	return c.gClient.GetRepo(nctx, req)
}

func (c *rClient) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	nctx := metadata.AppendToOutgoingContext(context.Background(), "auth-token", string(c.passkey))
	return c.gClient.ListFiles(nctx, req)
}

func (c *rClient) GetFile(ctx context.Context, req *pb.GetFileRequest) (*pb.GetFileResponse, error) {
	nctx := metadata.AppendToOutgoingContext(context.Background(), "auth-token", string(c.passkey))
	return c.gClient.GetFile(nctx, req)
}
func (c *rClient) UpdateFile(ctx context.Context, req *pb.UpdateFileRequest) (*pb.UpdateFileResponse, error) {
	nctx := metadata.AppendToOutgoingContext(context.Background(), "auth-token", string(c.passkey))
	return c.gClient.UpdateFile(nctx, req)
}
func (c *rClient) GetProjects(ctx context.Context, req *pb.GetProjectsRequest) (*pb.GetProjectsResponse, error) {
	nctx := metadata.AppendToOutgoingContext(context.Background(), "auth-token", string(c.passkey))
	return c.gClient.GetProjects(nctx, req)
}
func (c *rClient) GetReleases(ctx context.Context, req *pb.GetReleasesRequest) (*pb.GetReleasesResponse, error) {
	nctx := metadata.AppendToOutgoingContext(context.Background(), "auth-token", string(c.passkey))
	return c.gClient.GetReleases(nctx, req)
}
