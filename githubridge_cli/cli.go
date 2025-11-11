package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	pb "github.com/brotherlogic/githubridge/proto"
)

func printIssue(prefix string, issue *pb.GithubIssue) {
	fmt.Printf("%v%v\n", prefix, issue.GetTitle())
	for _, sissue := range issue.GetSubIssues() {
		printIssue(prefix+"  ", sissue)
	}
}

func main() {
	// Load the token
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	password, err := os.ReadFile(fmt.Sprintf("%v/.ghb", dirname))
	if err != nil {
		log.Fatalf("Can't read token: %v", err)
	}

	mContext := metadata.AppendToOutgoingContext(context.Background(), "auth-token", string(password))
	ctx, cancel := context.WithTimeout(mContext, time.Minute*60)
	defer cancel()

	conn, err := grpc.Dial(os.Args[1], grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Bad dial: %v", err)
	}

	client := pb.NewGithubridgeServiceClient(conn)

	switch os.Args[2] {
	case "issue":
		resp, err := client.GetIssue(ctx, &pb.GetIssueRequest{
			Repo: "gramophile",
			User: "brotherlogic",
			Id:   1760,
		})
		if err != nil {
			log.Fatalf("Unable to drain queue: %v", err)
		}
		printIssue("", &pb.GithubIssue{Title: resp.GetTitle(), SubIssues: resp.GetSubIssues()})
	case "issues":
		resp, err := client.GetIssues(ctx, &pb.GetIssuesRequest{})
		if err != nil {
			log.Fatalf("Unable to drain queue: %v", err)
		}
		for _, issue := range resp.GetIssues() {
			fmt.Printf("%v - %v / %v %v --> %v\n", issue.GetTitle(), issue.GetUser(), issue.GetRepo(), issue.GetId(), issue)
		}
	case "close":
		closeSet := flag.NewFlagSet("close", flag.ExitOnError)
		id := closeSet.Int64("id", -1, "Issue ID")
		repo := closeSet.String("repo", "", "Repo")
		user := closeSet.String("user", "brotherlogic", "User")
		if err := closeSet.Parse(os.Args[3:]); err == nil {
			_, err := client.CloseIssue(ctx, &pb.CloseIssueRequest{
				User: *user,
				Repo: *repo,
				Id:   *id,
			})
			if err != nil {
				log.Printf("Error on close: %v", err)
			}
		}
	case "open":
		openSet := flag.NewFlagSet("close", flag.ExitOnError)
		repo := openSet.String("repo", "", "Repo")
		user := openSet.String("user", "brotherlogic", "User")
		title := openSet.String("title", "", "Title")
		if err := openSet.Parse(os.Args[3:]); err == nil {
			_, err := client.CreateIssue(ctx, &pb.CreateIssueRequest{
				User:  *user,
				Repo:  *repo,
				Title: *title,
			})
			if err != nil {
				log.Printf("Error on close: %v", err)
			}
		}

	}
}
