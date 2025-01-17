package server

import (
	"context"

	pb "github.com/brotherlogic/githubridge/proto"
	"github.com/google/go-github/v50/github"
)

func (s *Server) GetProjects(ctx context.Context, req *pb.GetProjectsRequest) (*pb.GetProjectsResponse, error) {
	projects, _, err := s.client.Users.ListProjects(ctx, req.GetUser(), &github.ProjectListOptions{})
	if err != nil {
		return nil, err
	}

	var rprojects []string
	for _, project := range projects {
		rprojects = append(rprojects, project.GetName())
	}
	return &pb.GetProjectsResponse{Projects: rprojects}, nil
}
