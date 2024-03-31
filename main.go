package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/google/go-github/v50/github"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"

	pb "github.com/brotherlogic/githubridge/proto"
	"github.com/brotherlogic/githubridge/server"
)

var (
	port         = flag.Int("port", 8080, "Server port for grpc traffic")
	metricsPort  = flag.Int("metrics_port", 8081, "Metrics port")
	internalPort = flag.Int("internal_port", 8082, "Internal port")
	owner        = flag.String("owner", "brotherlogic", "")
)

func main() {
	flag.Parse()

	token := os.Getenv("GITHUBRIDGE_TOKEN")
	if token == "" {
		log.Fatalf("you must specify a valid github token under GITHUBRIDGE_TOKEN")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	s := server.NewServer(github.NewClient(tc), *owner)

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%v", *metricsPort), nil)
		log.Fatalf("gramophile is unable to serve metrics: %v", err)
	}()

	lis2, err := net.Listen("tcp", fmt.Sprintf(":%d", *internalPort))
	if err != nil {
		log.Fatalf("githubridge is unable to listen on the internal grpc port %v: %v", *internalPort, err)
	}
	gsInternal := grpc.NewServer()
	pb.RegisterGithubridgeServiceServer(gsInternal, s)

	go func() {
		if err := gsInternal.Serve(lis2); err != nil {
			log.Fatalf("githubridge is unable to serve internal grpc: %v", err)
		}
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("gramophile is unable to listen on the grpc port %v: %v", *port, err)
	}
	gs := grpc.NewServer(grpc.UnaryInterceptor(s.AuthCall))
	pb.RegisterGithubridgeServiceServer(gs, s)

	log.Printf("Serving on port :%d", *port)
	if err := gs.Serve(lis); err != nil {
		log.Fatalf("gramophile is unable to serve grpc: %v", err)
	}
	log.Fatalf("gramophile has closed the grpc port for some reason")
}
