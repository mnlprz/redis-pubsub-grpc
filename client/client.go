package main

import (
	"context"
	"ignite91/redis-pubsub-grpc/pubsub/pubsubpb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pubsubpb.NewPublisherClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Publish(ctx, &pubsubpb.PublishRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Response: ", r)
}
