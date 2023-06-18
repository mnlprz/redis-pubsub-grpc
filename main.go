package main

import (
	"context"
	"ignite91/redis-pubsub-grpc/pubsub/pubsubpb"
	"log"
	"net"

	"github.com/go-redis/redis"
	"google.golang.org/grpc"
)

type server struct {
	pubsubpb.UnimplementedPublisherServer
}

func (s *server) Publish(ctx context.Context, in *pubsubpb.PublishRequest) (*pubsubpb.PublishResponse, error) {
	redis := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	err := redis.Publish("send-user-data", "hola").Err()
	if err != nil {
		log.Panic(err)
	}
	log.Println("message published from GRPC: ")
	return &pubsubpb.PublishResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pubsubpb.RegisterPublisherServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
