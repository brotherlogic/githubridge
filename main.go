package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/google/go-github/v56/github"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	pb "github.com/brotherlogic/githubridge/proto"
)

var (
	port        = flag.Int("port", 8080, "Server port for grpc traffic")
	metricsPort = flag.Int("metrics_port", 8081, "Metrics port")
	owner       = flag.String("owner", "brotherlogic", "")
)

type Server struct {
	client *github.Client
	owner  string
}

func (s *Server) AddIssue(ctx context.Context, req *pb.AddIssueRequest) (*pb.AddIssueResponse, error) {
	issue, resp, err := s.client.Issues.Create(ctx, s.owner, req.GetJob(), &github.IssueRequest{
		Title: proto.String(req.GetTitle()),
		Body:  proto.String(req.GetBody()),
	})
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("Bad response code: %v", resp.StatusCode)
	}

	return &pb.AddIssueResponse{IssueId: int32(issue.GetID())}, nil
}

func main() {
	flag.Parse()

	token := os.Getenv("GITHUBRIDGE_TOKEN")
	if token == "" {
		log.Fatalf("you must specify a github token")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	s := &Server{
		owner:  *owner,
		client: github.NewClient(tc),
	}

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%v", *metricsPort), nil)
		log.Fatalf("gramophile is unable to serve metrics: %v", err)
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("gramophile is unable to listen on the grpc port %v: %v", *port, err)
	}
	gs := grpc.NewServer()
	pb.RegisterGithubBridegServer(gs, s)
	if err := gs.Serve(lis); err != nil {
		log.Fatalf("gramophile is unable to serve grpc: %v", err)
	}
	log.Fatalf("gramophile has closed the grpc port for some reason")
}
