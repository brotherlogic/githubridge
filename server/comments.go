package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/go-github/v50/github"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/brotherlogic/githubridge/proto"
)

func convertComment(comment *github.IssueComment) *pb.Comment {
	return &pb.Comment{
		Number: int32(*comment.ID),
		Text:   comment.GetBody(),
	}
}

func (s *Server) GetComments(ctx context.Context, req *pb.GetCommentsRequest) (*pb.GetCommentsResponse, error) {
	results, ghr, err := s.client.Issues.ListComments(ctx, req.GetUser(), req.GetRepo(), int(req.GetNumber()), &gituhb.IssueListCommentsOptions{})

	if err != nil {
		return nil, err
	}

	if ghr.LastPage != 1 {
		return nil, status.Errorf(codes.FailedPrecondition, "There are more comments than we're returning (%v)", req)
	}

	var comments []*pb.Comment
	for _, comment := range results {
		comments = append(comments, convertComment(comment))
	}
	return &pb.GetCommentsResponse{Comments: comments}, nil
}
