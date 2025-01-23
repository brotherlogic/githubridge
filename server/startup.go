package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	pb "github.com/brotherlogic/githubridge/proto"
	"github.com/google/go-github/v50/github"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/protobuf/proto"
)

var (
	trackedRepos = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "githubridge_tracked_repos",
		Help: "The number of repos being tracked",
	})

	trackedIssues = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "githubridge_tracked_issues",
		Help: "The number of issues being tracked",
	}, []string{"repo"})

	startupTime = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "githubridge_startup_time_ms",
	})

	callback = "http://ghwebhook.brotherlogic-backend.com/"
)

func convertIssue(issue *github.Issue, repo string) *pb.GithubIssue {
	log.Printf("CONVERT: %v", issue)

	issueState := pb.IssueState_ISSUE_STATE_UNKNOWN
	switch issue.GetState() {
	case "open":
		issueState = pb.IssueState_ISSUE_STATE_OPEN
	case "closed":
		issueState = pb.IssueState_ISSUE_STATE_CLOSED
	}

	var labels []string
	for _, label := range issue.Labels {
		labels = append(labels, label.GetName())
	}

	return &pb.GithubIssue{
		Id:         int64(issue.GetNumber()),
		Repo:       repo,
		User:       issue.GetUser().GetLogin(),
		Title:      issue.GetTitle(),
		OpenedDate: issue.GetCreatedAt().Unix(),
		State:      issueState,
		Labels:     labels,
	}
}

func (s *Server) loadIssues(ctx context.Context, repo string) error {
	cpage := 1
	lpage := 1

	for cpage <= lpage {
		// Read all the repos
		issues, resp, err := s.client.Issues.ListByRepo(ctx, s.user, repo, &github.IssueListByRepoOptions{
			ListOptions: github.ListOptions{Page: cpage},
		})
		lpage = resp.LastPage
		log.Printf("READ ISSUE: %v / %v (%v)", cpage, resp.LastPage, len(s.repos))
		if err != nil {
			return err
		}
		for _, issue := range issues {
			if !issue.IsPullRequest() {
				nissue := convertIssue(issue, repo)
				log.Printf("FOUND ISSUE: %v", nissue)
				s.issues = append(s.issues, nissue)
			}

		}

		cpage++
	}

	return nil
}

func (s *Server) startup(ctx context.Context) error {
	t := time.Now()
	defer func() {
		startupTime.Set(float64(time.Since(t).Milliseconds()))
	}()

	cpage := 1
	lpage := 1

	s.repos = []string{}
	s.issues = []*pb.GithubIssue{}

	// Install the webhook and health handler
	go func() {
		http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
			if s.ready {
				fmt.Fprintf(w, "OK")
			} else {
				fmt.Fprintf(w, "NO")
			}
		})
		http.HandleFunc("/", s.githubwebhook)

		err := http.ListenAndServe(fmt.Sprintf(":%v", 80), nil)
		if err != nil {
			panic(err)
		}
	}()

	for cpage <= lpage {
		// Read all the repos
		repos, resp, err := s.client.Repositories.List(ctx, "", &github.RepositoryListOptions{
			ListOptions: github.ListOptions{Page: cpage},
			Type:        "all",
		})
		lpage = resp.LastPage
		log.Printf("READ: %v / %v (%v)", cpage, resp.LastPage, len(s.repos))
		if err != nil {
			return err
		}
		for _, repo := range repos {
			log.Printf("REPO: %v", repo.GetName())
			s.repos = append(s.repos, repo.GetName())
		}

		cpage++

	}

	trackedRepos.Set(float64(len(s.repos)))

	// Ensure we have a callback on each repo
	for _, repo := range s.repos {
		s.loadIssues(ctx, repo)

		hooks, _, err := s.client.Repositories.ListHooks(ctx, s.user, repo, &github.ListOptions{})
		if err != nil {
			return err
		}

		events := []string{"issues", "issue_comment"}

		found := false
		foundAllEvents := true
		hid := int64(0)
		for _, h := range hooks {
			if h.Config["url"] == callback {
				hid = h.GetID()
				for _, event := range events {
					foundEvent := false
					for _, ex := range h.Events {
						if ex == event {
							foundEvent = true
						}
					}
					if !foundEvent {
						foundAllEvents = false
					}
				}
				found = true
			} else {
				log.Printf("Found %v", h.Config["url"])
			}
		}

		if !found || !foundAllEvents {
			log.Printf("Add to %v %v", hid, repo)
			hook := &github.Hook{
				Active: proto.Bool(true),
				Events: events,
				Config: map[string]interface{}{"url": callback},
			}
			if !found {
				a, b, c := s.client.Repositories.CreateHook(ctx, s.user, repo, hook)
				log.Printf("HERE: %v, %v, %v from %v", a, b, c, hook)
			} else {
				a, b, c := s.client.Repositories.EditHook(ctx, s.user, repo, hid, hook)
				log.Printf("HERE: %v, %v, %v from %v", a, b, c, hook)
			}
		}
	}

	return nil
}

func (s *Server) metrics() {
	issueMap := make(map[string]float64)
	for _, issue := range s.issues {
		issueMap[issue.GetRepo()]++
	}
	trackedIssues.Reset()
	for repo, count := range issueMap {
		trackedIssues.With(prometheus.Labels{"repo": repo}).Set(count)
	}
}
