package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	pb "github.com/brotherlogic/githubridge/proto"
	"github.com/google/go-github/v50/github"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
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
	case *github.IssuesEvent:
		repo := *event.Repo.Name
		action := *event.Action
		log.Printf("%v -> %v [%v]", repo, action, len(s.issues))
		if action == "closed" {
			issueCloses.Inc()
			var nissues []*pb.GithubIssue
			for _, issue := range s.issues {
				if issue.GetRepo() != repo ||
					issue.GetUser() != event.Repo.Owner.GetLogin() ||
					issue.GetId() != int64(event.Issue.GetNumber()) {
					nissues = append(nissues, issue)
				}
			}

			s.issues = nissues
			log.Printf("CLOSE: %v", len(s.issues))
			s.metrics()
		} else if action == "opened" {
			issueAdds.Inc()
			s.issues = append(s.issues, &pb.GithubIssue{
				Repo:       repo,
				User:       event.Repo.Owner.GetLogin(),
				Id:         int64(event.Issue.GetNumber()),
				Title:      event.Issue.GetTitle(),
				State:      pb.IssueState_ISSUE_STATE_OPEN,
				OpenedDate: time.Now().Unix(),
			})
			s.metrics()
		} else if action == "labeled" {
			for _, issue := range s.issues {
				if issue.GetRepo() == repo && issue.GetUser() == event.Repo.Owner.GetLogin() && issue.GetId() == int64(event.Issue.GetNumber()) {
					issue.Labels = append(issue.Labels, event.Label.GetName())
				}
			}
		}
	case *github.IssueCommentEvent:
		user := event.GetRepo().GetOwner().GetLogin()
		repo := event.GetRepo().GetName()
		isid := event.GetIssue().GetID()

		// Invalidate the cache
		key := fmt.Sprintf("%v-%v-%v", user, repo, isid)
		s.comments.Delete(key)
	case *github.LabelEvent:
		// Is this a label addition?
		user := event.GetRepo().GetOwner().GetLogin()
		repo := event.GetRepo().GetName()
		isid := event.GetLabel().GetID()

		for _, issue := range s.issues {
			if issue.GetRepo() == repo && issue.GetUser() == user && issue.GetId() == isid {
				issue.Labels = append(issue.Labels, event.GetLabel().GetName())
			}
		}
	default:
		log.Printf("Unable to process %v (%T)", event, event)
		w.Write([]byte(fmt.Sprintf("Unable to handle (%T)", event)))
	}

	log.Printf(" -> %v", len(s.issues))
}
