package server

import (
	"context"

	pb "github.com/brotherlogic/githubridge/proto"
	"github.com/google/go-github/v74/github"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetReleases(ctx context.Context, req *pb.GetReleasesRequest) (*pb.GetReleasesResponse, error) {
	releases, ghr, err := s.client.Repositories.ListReleases(ctx, req.GetUser(), req.GetRepo(), &github.ListOptions{})
	if err != nil {
		return nil, err
	}
	processResponse(ghr, "repos-listreleases")

	if ghr.LastPage != 1 {
		return nil, status.Errorf(codes.FailedPrecondition, "There are more comments than we're returning (%v)", req)
	}

	var versions []*pb.Version
	for _, release := range releases {
		versions = append(versions, &pb.Version{
			Version: *release.Name,
			Date:    release.CreatedAt.Unix(),
		})
	}
	return &pb.GetReleasesResponse{Versions: versions}, nil
}
