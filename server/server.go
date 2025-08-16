package server

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/google/go-github/v74/github"
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
	requests = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "githubridge_client_requests",
		Help: "The number of requests in a given type",
	}, []string{"type"})
)

func processResponse(resp *github.Response, typ string) {
	if resp != nil {
		quotaLeft.Set(float64(resp.Rate.Remaining))
		quotaAvail.Set(float64(resp.Rate.Limit))
		quotaResetTime.Set(float64(resp.Rate.Reset.UnixMilli()))
	}
	requests.With(prometheus.Labels{"type": typ}).Inc()
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
		client: client,
		user:   user,
		ready:  true}

	s.authKey = os.Getenv("GHB_AUTH_TOKEN")

	s.startup(context.Background())
	
	s.ready = true
	return s
}
