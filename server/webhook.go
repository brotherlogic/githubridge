package server

import (
	"io"
	"log"
	"net/http"

	"github.com/google/go-github/v50/github"
)

func (s *Server) githubwebhook(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Unable to parse body: %v", err)
		return
	}

	event, err := github.ParseWebHook(github.WebHookType(r), body)
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
