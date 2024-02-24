package server

import (
	"fmt"
	"log"
	"net/http"

	pb "github.com/brotherlogic/githubridge/proto"
	"github.com/google/go-github/v50/github"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/protobuf/proto"
)

var (
	issueAdds = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "githubridge_issue_adds",
		Help: "The number of repos being tracked",
	})
	issueCloses = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "githubridge_issue_closes",
		Help: "The number of repos being tracked",
	})
	pings = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "githubridge_pings",
		Help: "The number of repos being tracked",
	}, []string{"type"})
	webhooks = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "githubridge_webhooks",
		Help: "The number of repos being tracked",
	})
)

func (s *Server) githubwebhook(w http.ResponseWriter, r *http.Request) {
	webhooks.Inc()
	payload, err := github.ValidatePayload(r, nil)
	if err != nil {
		log.Printf("Bad payload: %v", err)
	}

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		log.Printf("Bad stuff: %v", err)
	}
	pings.With(prometheus.Labels{"type": fmt.Sprintf("%T", event)})

	switch event := event.(type) {
	case *github.IssueEvent:
		repo := event.Issue.Repository.Name
		action := event.Event
		log.Printf("%v -> %v", repo, action)
		if action == proto.String("closed") {
			issueCloses.Inc()
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
		} else if action == proto.String("open") {
			issueAdds.Inc()
			s.issues = append(s.issues, &pb.GithubIssue{
				Repo:  *repo,
				User:  event.Issue.Repository.Owner.GetName(),
				Id:    event.Issue.GetID(),
				Title: event.Issue.GetTitle(),
			})
		}
	default:
		log.Printf("Unable to process %v (%T)", event, event)
	}
}
