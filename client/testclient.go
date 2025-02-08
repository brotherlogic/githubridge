package githubridgeclient

import (
	"context"
	"fmt"

	pb "github.com/brotherlogic/githubridge/proto"
)

type TestClient struct {
	counter  int64
	issues   []*pb.GithubIssue
	labels   map[string][]string
	comments map[string][]*pb.Comment
}

func GetTestClient() GithubridgeClient {
	return &TestClient{
		issues:   []*pb.GithubIssue{},
		counter:  0,
		labels:   make(map[string][]string),
		comments: make(map[string][]*pb.Comment),
	}
}

func (c *TestClient) AddLabel(ctx context.Context, req *pb.AddLabelRequest) (*pb.AddLabelResponse, error) {
	label := fmt.Sprintf("%v-%v-%v", req.GetUser(), req.GetRepo(), req.GetId())
	c.labels[label] = append(c.labels[label], req.GetLabel())
	return &pb.AddLabelResponse{}, nil
}

func (c *TestClient) CreateIssue(ctx context.Context, req *pb.CreateIssueRequest) (*pb.CreateIssueResponse, error) {
	c.counter++
	c.issues = append(c.issues, &pb.GithubIssue{Id: c.counter, Title: req.GetTitle(), Repo: req.GetRepo()})
	return &pb.CreateIssueResponse{IssueId: c.counter}, nil
}

func (c *TestClient) CloseIssue(ctx context.Context, req *pb.CloseIssueRequest) (*pb.CloseIssueResponse, error) {
	var nissues []*pb.GithubIssue
	for _, issue := range c.issues {
		if issue.GetRepo() != req.GetRepo() || issue.GetId() != req.GetId() {
			nissues = append(nissues, issue)
		}
	}
	c.issues = nissues
	return &pb.CloseIssueResponse{}, nil
}

func (c *TestClient) CommentOnIssue(ctx context.Context, req *pb.CommentOnIssueRequest) (*pb.CommentOnIssueResponse, error) {
	return &pb.CommentOnIssueResponse{}, nil
}

func (c *TestClient) GetIssue(ctx context.Context, req *pb.GetIssueRequest) (*pb.GetIssueResponse, error) {
	return &pb.GetIssueResponse{}, nil
}

func (c *TestClient) GetRepos(ctx context.Context, req *pb.GetReposRequest) (*pb.GetReposResponse, error) {
	return &pb.GetReposResponse{}, nil
}

func (c *TestClient) GetRepo(ctx context.Context, req *pb.GetRepoRequest) (*pb.GetRepoResponse, error) {
	return &pb.GetRepoResponse{}, nil
}

func (c *TestClient) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	return &pb.ListFilesResponse{}, nil
}

func (c *TestClient) GetFile(ctx context.Context, req *pb.GetFileRequest) (*pb.GetFileResponse, error) {
	return &pb.GetFileResponse{}, nil
}

func (c *TestClient) UpdateFile(ctx context.Context, req *pb.UpdateFileRequest) (*pb.UpdateFileResponse, error) {
	return &pb.UpdateFileResponse{}, nil
}

func (c *TestClient) GetLabels(ctx context.Context, req *pb.GetLabelsRequest) (*pb.GetLabelsResponse, error) {
	return &pb.GetLabelsResponse{Labels: c.labels[fmt.Sprintf("%v-%v-%v", req.GetUser(), req.GetRepo(), req.GetId())]}, nil
}

func (c *TestClient) GetIssues(ctx context.Context, req *pb.GetIssuesRequest) (*pb.GetIssuesResponse, error) {
	return &pb.GetIssuesResponse{Issues: c.issues}, nil
}

func (c *TestClient) GetComments(ctx context.Context, req *pb.GetCommentsRequest) (*pb.GetCommentsResponse, error) {
	return &pb.GetCommentsResponse{Comments: c.comments[fmt.Sprintf("%v-%v-%v", req.GetUser(), req.GetRepo(), req.GetId())]}, nil
}

func (c *TestClient) DeleteLabel(ctx context.Context, req *pb.DeleteLabelRequest) (*pb.DeleteLabelResponse, error) {
	label := fmt.Sprintf("%v-%v-%v", req.GetUser(), req.GetRepo(), req.GetId())
	delete(c.labels, label)
	return &pb.DeleteLabelResponse{}, nil
}
