package server

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/google/go-github/v50/github"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/sync/syncmap"

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

type CommentCache struct {
	Comments []*pb.Comment
	Cached   time.Time
}

type Server struct {
	client  *github.Client
	repos   []string
	user    string
	issues  []*pb.GithubIssue
	ready   bool // ready to server
	authKey string

	comments syncmap.Map
}

func NewServer(client *github.Client, user string) *Server {
	s := &Server{
		client:   client,
		comments: map[string]*CommentCache{},
		user:     user,
		ready:    true}

	s.authKey = os.Getenv("GHB_AUTH_TOKEN")

	err := s.startup(context.Background())
	if err != nil {
		log.Printf("Failed to startup: %v", err)
		os.Exit(1)
	}

	s.ready = true
	return s
}
