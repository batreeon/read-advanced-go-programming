package main

import (
	"context"
	"fmt"
	"io"
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

	stream, err := client.Subscribe(context.Background(), &pubsub.String{Value: "golang:"})
	if err != nil {
		log.Fatal(err)
	}
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}
}
