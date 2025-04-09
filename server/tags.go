package server

import (
	"context"

	pb "github.com/brotherlogic/githubridge/proto"
	"github.com/google/go-github/v50/github"
	"google.golang.org/protobuf/proto"
)

func (s *Server) GetTags(ctx context.Context, req *pb.GetTagsRequest) (*pb.GetTagsResponse, error) {
	refs, resp, err := s.client.Git.ListMatchingRefs(ctx, req.GetUser(), req.GetRepo(), &github.ReferenceListOptions{
		Ref: "tags",
	})
	if err != nil {
		return nil, err
	}
	processResponse(resp, "get_tags")

	var tags []string
	for _, val := range refs {
		tags = append(tags, val.GetRef())
	}
	return &pb.GetTagsResponse{Tags: tags}, nil
}

func (s *Server) SetTag(ctx context.Context, req *pb.SetTagRequest) (*pb.SetTagResponse, error) {
	_, resp, err := s.client.Git.CreateRef(ctx, req.GetUser(), req.GetRepo(), &github.Reference{
		Ref: proto.String(req.GetTag()),
	})
	if err != nil {
		return nil, err
	}
	processResponse(resp, "set_tag")

	return &pb.SetTagResponse{}, nil
}
