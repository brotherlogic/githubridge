package server

import (
	"context"

	pb "github.com/brotherlogic/githubridge/proto"
)

func (s *Server) GetRepos(ctx context.Context, req *pb.GetReposRequest) (*pb.GetReposResponse, error) {
	return &pb.GetReposResponse{Repos: s.repos}, nil
}
