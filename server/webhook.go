package server

import (
	"log"
	"net/http"

	pb "github.com/brotherlogic/githubridge/proto"
	"github.com/google/go-github/v50/github"
	"google.golang.org/protobuf/proto"
)

func (s *Server) githubwebhook(w http.ResponseWriter, r *http.Request) {
	payload, err := github.ValidatePayload(r, nil)

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		log.Printf("Bad stuff: %v", err)
	}
	switch event := event.(type) {
	case *github.IssueEvent:
		repo := event.Issue.Repository.Name
		action := event.Event
		log.Printf("%v -> %v", repo, action)
		if action == proto.String("closed") {
			var nissues []*pb.GithubIssue
			for _, issue := range s.issues {
				if issue.GetRepo() != *repo &&
					issue.GetUser() != *event.Issue.Repository.Owner.Name &&
					issue.GetId() != event.Issue.GetID() {
					nissues = append(nissues, issue)
				}
			}

			s.issues = nissues
			trackedIssues.Set(float64(len(s.issues)))
		}
	default:
		log.Printf("Unable to process %v (%T)", event, event)
	}
}
