package server

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/go-github/v50/github"

	pb "github.com/brotherlogic/githubridge/proto"
)

func convertComment(comment *github.IssueComment) *pb.Comment {
	return &pb.Comment{
		Number: int32(*comment.ID),
		Text:   comment.GetBody(),
	}
}

func (s *Server) getFromCommentCache(ctx context.Context, req *pb.GetCommentsRequest) ([]*pb.Comment, error) {
	key := fmt.Sprintf("%v-%v-%v", req.GetUser(), req.GetRepo(), req.GetId())
	val, ok := s.comments[key]
	if ok && time.Since(val.Cached) < time.Minute*30 {
		return val.Comments, nil
	}
	return nil, status.Errorf(codes.NotFound, "Not in cache")
}

func (s *Server) insertCommentsIntoCache(ctx context.Context, req *pb.GetCommentsRequest, comments []*pb.Comment) error {
	key := fmt.Sprintf("%v-%v-%v", req.GetUser(), req.GetRepo(), req.GetId())
	s.comments[key] = &CommentCache{
		Comments: comments,
		Cached:   time.Now(),
	}
	return nil
}

func (s *Server) GetComments(ctx context.Context, req *pb.GetCommentsRequest) (*pb.GetCommentsResponse, error) {
	ccomments, err := s.getFromCommentCache(ctx, req)
	if err == nil {
		return &pb.GetCommentsResponse{Comments: ccomments}, nil
	}

	results, ghr, err := s.client.Issues.ListComments(ctx, req.GetUser(), req.GetRepo(), int(req.GetId()), &github.IssueListCommentsOptions{})
	processResponse(ghr)

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
	s.insertCommentsIntoCache(ctx, req, comments)
	return &pb.GetCommentsResponse{Comments: comments}, nil
}
