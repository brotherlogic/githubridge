package server

import (
	"context"

	pb "github.com/brotherlogic/githubridge/proto"

	"github.com/google/go-github/v50/github"
)

func (s *Server) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	path := req.GetPath()
	if path == "" {
		path = "/"
	}
	_, content, ghr, err := s.client.Repositories.GetContents(ctx, req.GetUser(), req.GetRepo(), path, &github.RepositoryContentGetOptions{})

	processResponse(ghr)

	if err != nil {
		return nil, err
	}

	var files []*pb.File
	for _, c := range content {
		files = append(files, &pb.File{
			Name: c.GetPath(),
			Hash: c.GetSHA(),
		})
	}

	return &pb.ListFilesResponse{Files: files}, nil
}
