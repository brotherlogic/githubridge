package server

import (
	"context"

	pb "github.com/brotherlogic/githubridge/proto"
	"google.golang.org/protobuf/proto"

	"github.com/google/go-github/v50/github"
)

func (s *Server) GetTags(ctx context.Context, req *pb.GetTagsRequest) (*pb.GetTagsResponse, error) {
	tags, ghr, err := s.client.Repositories.ListTags(ctx, req.GetUser(), req.GetRepo(), &github.ListOptions{
		PerPage: 100,
	})
	if err != nil {
		return nil, err
	}
	processResponse(ghr, "repos-listtags")

	var rtags []*pb.Tag
	for _, tag := range tags {
		rtags = append(rtags, &pb.Tag{
			Value:       tag.GetName(),
			TimestampMs: tag.GetCommit().GetCommitter().GetDate().UnixMilli(),
			Sha5:        tag.GetCommit().GetSHA(),
		})
	}

	return &pb.GetTagsResponse{Tags: rtags}, nil
}

func (s *Server) SetTag(ctx context.Context, req *pb.SetTagRequest) (*pb.SetTagResponse, error) {
	_, ghr, err := s.client.Git.CreateTag(ctx, req.GetUser(), req.GetRepo(), &github.Tag{
		Tag: proto.String(req.GetTag()),
		SHA: proto.String(req.GetSha5()),
	})
	if err != nil {
		return nil, err
	}
	processResponse(ghr, "git-createtag")

	return &pb.SetTagResponse{}, nil
}
