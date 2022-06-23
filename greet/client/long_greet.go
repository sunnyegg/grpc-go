package main

import (
	"context"
	"log"
	"time"

	pb "github.com/sunnyegg/grpc-go/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	req := []*pb.GreetRequest{
		{FirstName: "Sunny"},
		{FirstName: "Adila"},
		{FirstName: "Tresna"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, r := range req {
		log.Printf("Sending req: %v\n", r)
		stream.Send(r)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %v\n", res.Result)
}
