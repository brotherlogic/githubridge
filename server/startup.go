package server

import (
	"context"
	"log"

	"github.com/google/go-github/v50/github"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	trackedRepos = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "githubridge_tracked_repos",
		Help: "The number of repos being tracked",
	})
)

func (s *Server) startup(ctx context.Context) error {
	cpage := 1
	lpage := 1

	for cpage <= lpage {
		// Read all the repos
		repos, resp, err := s.client.Repositories.List(ctx, s.user, &github.RepositoryListOptions{
			ListOptions: github.ListOptions{Page: cpage},
		})
		lpage = resp.LastPage
		log.Printf("READ: %v", resp)
		if err != nil {
			return err
		}

		s.repos = []string{}
		for _, repo := range repos {
			s.repos = append(s.repos, repo.GetName())
		}

		cpage++

	}

	trackedRepos.Set(float64(len(s.repos)))

	return nil
}
