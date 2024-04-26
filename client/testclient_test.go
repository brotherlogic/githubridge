package githubridgeclient

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/githubridge/proto"
)

func TestAddLabel(t *testing.T) {
	s := GetTestClient()

	s.AddLabel(context.Background(), &pb.AddLabelRequest{
		User:  "user",
		Repo:  "repo",
		Id:    1,
		Label: "donkey",
	})

	labels, err := s.GetLabels(context.Background(), &pb.GetLabelsRequest{
		User: "user",
		Repo: "repo",
		Id:   1,
	})

	if err != nil {
		t.Fatalf("Error getting label: %v", err)
	}

	if len(labels.GetLabels()) != 1 || labels.GetLabels()[0] != "donkey" {
		t.Errorf("Bad labels: %v", labels)
	}
}
