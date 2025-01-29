package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/google/go-github/v66/github"

	ghiter "github.com/enrichman/gh-iter"

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
	val := valo.(*CommentCache)
	if ok && time.Since(val.Cached) < time.Minute*30 {
		return val.Comments, nil
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
	log.Printf("GetComments %v", req)

	ccomments, err := s.getFromCommentCache(ctx, req)
	if err == nil {
		log.Printf("Serving from cache")
		return &pb.GetCommentsResponse{Comments: ccomments}, nil
	}

	// create an iterator and start looping through all the results
	rcomments := ghiter.NewFromFn1(s.client.Issues.ListComments, req.GetUser(), req.GetRepo(), int(req.GetId()))
	var comments []*pb.Comment
	for _, rcomment := range rcomments {
		comments = append(comments, convertComment(rcomment))
	}
	if err := rcomments.Err(); err != nil {
		return nil, err
	}

	s.insertCommentsIntoCache(ctx, req, comments)
	return &pb.GetCommentsResponse{Comments: comments}, nil
}

func (s *Server) CommentOnIssue(ctx context.Context, req *pb.CommentOnIssueRequest) (*pb.CommentOnIssueResponse, error) {
	_, resp, err := s.client.Issues.CreateComment(ctx, req.GetUser(), req.GetRepo(), int(req.GetId()), &github.IssueComment{
		Body: proto.String(req.GetComment()),
	})
	processResponse(resp)

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("bad response code for comment: %v", resp.StatusCode)
	}

	return &pb.CommentOnIssueResponse{}, nil
}
