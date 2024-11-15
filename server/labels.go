package server

import (
	"context"

	pb "github.com/brotherlogic/githubridge/proto"
	"github.com/google/go-github/v50/github"
)

func (s *Server) GetLabels(ctx context.Context, req *pb.GetLabelsRequest) (*pb.GetLabelsResponse, error) {
	r, gr, err := s.client.Issues.ListLabelsByIssue(ctx, req.GetUser(), req.GetRepo(), int(req.GetId()), &github.ListOptions{})
	if err != nil {
		return nil, err
	}
	processResponse(gr)

	var labels []string
	for _, label := range r {
		labels = append(labels, label.GetName())
	}
	return &pb.GetLabelsResponse{Labels: labels}, nil
}

func (s *Server) AddLabel(ctx context.Context, req *pb.AddLabelRequest) (*pb.AddLabelResponse, error) {
	_, _, err := s.client.Issues.AddLabelsToIssue(ctx, req.GetUser(), req.GetRepo(), int(req.GetId()), []string{req.GetLabel()})

	return &pb.AddLabelResponse{}, err
}

func (s *Server) DeleteLabel(ctx context.Context, req *pb.DeleteLabelRequest) (*pb.DeleteLabelResponse, error) {
	_, err := s.client.Issues.RemoveLabelForIssue(ctx, req.GetUser(), req.GetRepo(), int(req.GetId()), req.GetLabel())

	return &pb.DeleteLabelResponse{}, err
}
