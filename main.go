package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/google/go-github/v74/github"
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

	lisInt, err := net.Listen("tcp", fmt.Sprintf(":%d", *internalPort))
	if err != nil {
		log.Fatalf("mdb is unable to listen on the internal grpc port %v: %v", *port, err)
	}
	gsInt := grpc.NewServer(grpc.ChainUnaryInterceptor(
		s.ServerInterceptor,
	))
	pb.RegisterGithubridgeServiceServer(gsInt, s)

	go func() {
		err := gsInt.Serve(lisInt)
		log.Fatalf("mdb is unable to sever grpc internally: %v", err)
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("gramophile is unable to listen on the grpc port %v: %v", *port, err)
	}
	gs := grpc.NewServer(
		grpc.UnaryInterceptor(s.AuthCall), grpc.ChainUnaryInterceptor(s.ServerInterceptor))
	pb.RegisterGithubridgeServiceServer(gs, s)

	log.Printf("Serving on port :%d", *port)
	if err := gs.Serve(lis); err != nil {
		log.Fatalf("gramophile is unable to serve grpc: %v", err)
	}
	log.Fatalf("gramophile has closed the grpc port for some reason")
}
