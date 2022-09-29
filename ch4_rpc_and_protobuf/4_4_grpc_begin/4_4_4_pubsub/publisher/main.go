package main

import (
	"context"
	"log"
	"4_4_4_pubsub/pubsub"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pubsub.NewPubsubServiceClient(conn)
	
	_, err = client.Publish(
		context.Background(), &pubsub.String{Value: "golang: hello Go"},
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Publish(
		context.Background(), &pubsub.String{Value: "docker: hello Docker"},
	)
	if err != nil {
		log.Fatal(err)
	}
}
