package server

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/google/go-github/v74/github"

	pb "github.com/brotherlogic/githubridge/proto"
)

func convertComment(comment *github.IssueComment) *pb.Comment {
	return &pb.Comment{
		Number:    int32(*comment.ID),
		Text:      comment.GetBody(),
		Timestamp: comment.CreatedAt.Unix(),
	}
}

func (s *Server) getFromCommentCache(ctx context.Context, req *pb.GetCommentsRequest) ([]*pb.Comment, error) {
	key := fmt.Sprintf("%v-%v-%v", req.GetUser(), req.GetRepo(), req.GetId())
	valo, ok := s.comments.Load(key)

	if ok {
		val := valo.(*CommentCache)
		if time.Since(val.Cached) < time.Minute*30 {
			return val.Comments, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "Not in cache")
}

func (s *Server) insertCommentsIntoCache(ctx context.Context, req *pb.GetCommentsRequest, comments []*pb.Comment) error {
	key := fmt.Sprintf("%v-%v-%v", req.GetUser(), req.GetRepo(), req.GetId())
	s.comments.Store(key, &CommentCache{
		Comments: comments,
		Cached:   time.Now(),
	})
	return nil
}

func (s *Server) GetComments(ctx context.Context, req *pb.GetCommentsRequest) (*pb.GetCommentsResponse, error) {
	ccomments, err := s.getFromCommentCache(ctx, req)
	if err == nil {
		return &pb.GetCommentsResponse{Comments: ccomments}, nil
	}

	results, ghr, err := s.client.Issues.ListComments(ctx, req.GetUser(), req.GetRepo(), int(req.GetId()), &github.IssueListCommentsOptions{})
	processResponse(ghr, "issues-listcomments")

	if err != nil {
		return nil, err
	}

	if ghr.LastPage > 1 {
		return nil, status.Errorf(codes.FailedPrecondition, "There are more comments than we're returning (%v)  -> %+v", req, ghr)
	}

	var comments []*pb.Comment
	for _, comment := range results {
		comments = append(comments, convertComment(comment))
	}
	s.insertCommentsIntoCache(ctx, req, comments)
	return &pb.GetCommentsResponse{Comments: comments}, nil
}

func (s *Server) CommentOnIssue(ctx context.Context, req *pb.CommentOnIssueRequest) (*pb.CommentOnIssueResponse, error) {
	_, resp, err := s.client.Issues.CreateComment(ctx, req.GetUser(), req.GetRepo(), int(req.GetId()), &github.IssueComment{
		Body: proto.String(req.GetComment()),
	})
	processResponse(resp, "issues-createcomment")

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("bad response code for comment: %v", resp.StatusCode)
	}

	return &pb.CommentOnIssueResponse{}, nil
}
