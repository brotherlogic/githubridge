package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/brotherlogic/githubridge/proto"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*60)
	defer cancel()

	conn, err := grpc.Dial(os.Args[1], grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Bad dial: %v", err)
	}

	conn, serr := grpc.Dial(os.Args[1], grpc.WithTransportCredentials(insecure.NewCredentials()))
	if serr != nil {
		log.Fatalf("Bad dial: %v", err)
	}

	client := pb.NewGithubridgeServiceClient(conn)

	switch os.Args[2] {
	case "issues":
		resp, err := client.GetIssues(ctx, &pb.GetIssuesRequest{})
		if err != nil {
			log.Fatalf("Unable to drain queue: %v", err)
		}
		for _, issue := range resp.GetIssues() {
			fmt.Printf("Drained %v items\n", issue.GetTitle())
		}
	}
}