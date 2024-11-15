package server

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/v50/github"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	pb "github.com/brotherlogic/githubridge/proto"
)

var (
	quotaLeft = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "githubridge_quota_left",
		Help: "The amount of quota left",
	})
	quotaAvail = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "githubridge_quota_available",
		Help: "The amount of quota left",
	})
	quotaResetTime = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "githubridge_quota_reset_time",
		Help: "The amount of quota left",
	})
)

func processResponse(resp *github.Response) {
	quotaLeft.Set(float64(resp.Rate.Remaining))
	quotaAvail.Set(float64(resp.Rate.Limit))
	quotaResetTime.Set(float64(resp.Rate.Reset.UnixMilli()))
}

type Server struct {
	client  *github.Client
	repos   []string
	user    string
	issues  []*pb.GithubIssue
	ready   bool // ready to server
	authKey string
}

func NewServer(client *github.Client, user string) *Server {
	s := &Server{client: client, user: user, ready: true}
	err := s.startup(context.Background())
	if err != nil {
		log.Printf("Failed to startup: %v", err)
		os.Exit(1)
	}

	s.authKey = os.Getenv("GHB_AUTH_TOKEN")
	log.Printf("Loaded with %v", s.authKey)

	s.ready = true
	return s
}
