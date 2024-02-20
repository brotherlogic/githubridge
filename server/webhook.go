package server

import (
	"log"
	"net/http"

	"github.com/google/go-github/v50/github"
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
	default:
		log.Printf("Unable to process %v (%T)", event, event)
	}
}
