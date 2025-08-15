package server

import (
	"context"
	"io"
	"log"

	pb "github.com/brotherlogic/githubridge/proto"
	"google.golang.org/protobuf/proto"

	"github.com/google/go-github/v74/github"
)

func (s *Server) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	log.Printf("ListFiles: %v", req)
	path := req.GetPath()
	if path == "" {
		path = "/"
	}
	_, content, ghr, err := s.client.Repositories.GetContents(ctx, req.GetUser(), req.GetRepo(), path, &github.RepositoryContentGetOptions{})
	log.Printf("Path: %v -> %v err, with %+v", path, err, ghr)

	processResponse(ghr, "repos-getcontents")

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

func (s *Server) GetFile(ctx context.Context, req *pb.GetFileRequest) (*pb.GetFileResponse, error) {
	rcs, ghr, err := s.client.Repositories.DownloadContents(ctx, req.GetUser(), req.GetRepo(), req.GetPath(), &github.RepositoryContentGetOptions{})

	processResponse(ghr, "repos-downloadcontents")

	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(rcs)
	if err != nil {
		return nil, err
	}
	rcs.Close()

	return &pb.GetFileResponse{Content: bodyBytes}, nil
}

func (s *Server) UpdateFile(ctx context.Context, req *pb.UpdateFileRequest) (*pb.UpdateFileResponse, error) {
	_, ghr, err := s.client.Repositories.UpdateFile(ctx, req.GetUser(), req.GetRepo(), req.GetPath(),
		&github.RepositoryContentFileOptions{
			Message: proto.String(req.GetMessage()),
			Content: req.GetContent(),
		})
	processResponse(ghr, "repos-updatefile")
	return &pb.UpdateFileResponse{}, err
}
