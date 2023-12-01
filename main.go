package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"

	"github.com/brotherlogic/githubridge/server"
	"github.com/google/go-github/v50/github"

	pb "github.com/brotherlogic/githubridge/proto"
)

var (
	port = flag.Int("port", 8080, "Server port for grpc traffic")
)

func main() {
	accessCode := os.Getenv("GITHUB_COD")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessCode},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	s := server.New(github.NewClient(tc))

	gs := grpc.NewServer()
	pb.RegisterGithubridgeServer(gs, s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("githubridge is unable to listen on the grpc port %v: %v", *port, err)
	}
	if err := gs.Serve(lis); err != nil {
		log.Fatalf("githubridge is unable to serve grpc: %v", err)
	}
	log.Fatalf("githubridge has closed the grpc port for some reason")
}
