package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

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

	callback = "http://ghwebhook.brotherlogic-backend.com/"
)

func (s *Server) startup(ctx context.Context) error {
	cpage := 1
	lpage := 1

	s.repos = []string{}
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
				PingURL: proto.String(callback),
				URL:     proto.String(callback),
				TestURL: proto.String(callback),
				Name:    proto.String("KubeHookv1"),
				Events:  []string{"issues"},
				Type:    proto.String("json"),
			}
			a, b, c := s.client.Repositories.CreateHook(ctx, s.user, repo, hook)
			log.Printf("HERE: %v, %v, %v from %v", a, b, c, hook)
		}
	}

	// Install the webhook
	http.HandleFunc("/githubwebhook", s.githubwebhook)
	err := http.ListenAndServe(fmt.Sprintf(":%v", 80), nil)
	if err != nil {
		panic(err)
	}

	return nil
}
