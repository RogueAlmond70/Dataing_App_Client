package main

import (
	"context"
	"log"
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

	likedYou := c.ListLikedYou(ctx)
}
