package server

import (
	"context"

	pb "github.com/brotherlogic/githubridge/proto"
)

func (s *Server) GetRepos(ctx context.Context, req *pb.GetReposRequest) (*pb.GetReposResponse, error) {
	return &pb.GetReposResponse{Repos: s.repos}, nil
}

func (s *Server) GetRepo(ctx context.Context, req *pb.GetRepoRequest) (*pb.GetRepoResponse, error) {
	branch := req.GetBranch()
	if branch == "" {
		branch = "main"
	}
	repo, ghr, err := s.client.Repositories.GetBranch(ctx, req.GetUser(), req.GetRepo(), branch, true)
	processResponse(ghr)
	if err != nil {
		return nil, err
	}

	return &pb.GetRepoResponse{Sha1: repo.GetCommit().GetSHA()}, nil
}
