package server

import (
	"context"

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
	// Read all the repos
	repos, _, err := s.client.Repositories.List(ctx, s.user, &github.RepositoryListOptions{})
	if err != nil {
		return err
	}

	s.repos = []string{}
	for _, repo := range repos {
		s.repos = append(s.repos, repo.GetName())
	}

	trackedRepos.Set(float64(len(s.repos)))

	return nil
}
