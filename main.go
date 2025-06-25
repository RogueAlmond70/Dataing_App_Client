package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/RogueAlmond70/Muzz/explorepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials())) // Change this to secure with TLS
	if err != nil {

		log.Fatalf("failed to connect to grpc server: %v", err)
	}
	defer conn.Close()
	c := pb.NewExploreServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second) // Times out after a second.
	defer cancel()

	likedYouRequest := pb.ListLikedYouRequest{
		RecipientUserId: "1",
		//pagination_token: "2",   This is optional, but look into it.
	}

	newLikedYouRequest := pb.ListLikedYouRequest{ // Same request type, look into this.
		RecipientUserId: "1",
		//pagination_token: "2",   This is optional, but look into it.
	}

	countLikedYouRequest := pb.CountLikedYouRequest{
		RecipientUserId: "1",
	}

	putdecisionrequest := pb.PutDecisionRequest{
		ActorUserId:     "1",
		RecipientUserId: "2",
		LikedRecipient:  false,
	}

	likedYou, err := c.ListLikedYou(ctx, &likedYouRequest)
	if err != nil {
		fmt.Errorf("error sending liked you request: %v", err)
		os.Exit(1)
	}
	for _, y := range likedYou.Likers {
		fmt.Println(y.GetActorId())
	}

	newlikedYou, err := c.ListNewLikedYou(ctx, &newLikedYouRequest)
	if err != nil {
		fmt.Errorf("error sending new liked you request: %v", err)
		os.Exit(1)
	}
	for _, y := range newlikedYou.Likers {
		fmt.Println(y.GetActorId())
	}

	countLikedYou, err := c.CountLikedYou(ctx, &countLikedYouRequest)
	if err != nil {
		fmt.Errorf("error sending count liked you request: %v", err)
		os.Exit(1)
	}
	fmt.Println("You have %v likes!", countLikedYou.Count)

	putDecisionResponse, err := c.PutDecision(ctx, &putdecisionrequest)
	if err != nil {
		fmt.Errorf("error sending out decision request: %v", err) // Log this err and the other errors.
		os.Exit(1)
	}

	mutualMatch := putDecisionResponse.GetMutualLikes()
	fmt.Println("Mutual match with user: %v", mutualMatch)

}
