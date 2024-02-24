package server

import (
	"context"
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
	trackedRepos = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "githubridge_tracked_repos",
		Help: "The number of repos being tracked",
	})

	trackedIssues = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "githubridge_tracked_issues",
		Help: "The number of issues being tracked",
	})

	callback = "http://ghwebhook.brotherlogic-backend.com/"
)

func convertIssue(issue *github.Issue) *pb.GithubIssue {
	log.Printf("CONVERT: %v", issue)
	return &pb.GithubIssue{
		Id:    issue.GetID(),
		Repo:  issue.GetRepository().GetName(),
		User:  issue.GetRepository().GetOwner().GetLogin(),
		Title: issue.GetTitle(),
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
				nissue := convertIssue(issue)
				log.Printf("FOUND ISSUE: %v", nissue)
				s.issues = append(s.issues, nissue)
			}

		}

		cpage++
	}

	return nil
}

func (s *Server) startup(ctx context.Context) error {
	cpage := 1
	lpage := 1

	s.repos = []string{}
	s.issues = []*pb.GithubIssue{}

	// Install the webhook
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
		repos, resp, err := s.client.Repositories.List(ctx, s.user, &github.RepositoryListOptions{
			ListOptions: github.ListOptions{Page: cpage},
		})
		lpage = resp.LastPage
		log.Printf("READ: %v / %v (%v)", cpage, resp.LastPage, len(s.repos))
		if err != nil {
			return err
		}
		for _, repo := range repos {
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

		found := false
		for _, h := range hooks {
			if h.Config["url"] == callback {
				found = true
			} else {
				log.Printf("Found %v", h.Config["url"])
			}
		}

		if !found {
			log.Printf("Add to %v", repo)
			hook := &github.Hook{
				Active: proto.Bool(true),
				Events: []string{"issues"},
				Config: map[string]interface{}{"url": callback},
			}
			a, b, c := s.client.Repositories.CreateHook(ctx, s.user, repo, hook)
			log.Printf("HERE: %v, %v, %v from %v", a, b, c, hook)
		}
	}

	trackedIssues.Set(float64(len(s.issues)))

	return nil
}
