package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/brotherlogic/githubridge/proto"

	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, err := grpc.Dial(os.Args[1], grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Bad dial: %v", err)
	}

	client := pb.NewGithubBridegClient(conn)

	switch os.Args[2] {
	case "add":
		res, err := client.AddIssue(ctx, &pb.AddIssueRequest{
			Title: "Testing",
			Job:   "githubridge",
		})
		fmt.Printf("%v -> %v\n", res, err)
	default:
		fmt.Printf("Unknown command: %v\n", os.Args[1])
	}
}
