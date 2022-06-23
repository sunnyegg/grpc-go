package main

import (
	"context"
	"io"
	"log"

	pb "github.com/sunnyegg/grpc-go/calculator/proto"
)

func doPrime(c pb.CalculatorServiceClient, number int32) {
	log.Println("doPrime was invoked")

	req := &pb.PrimeRequest{
		Number: number,
	}

	stream, err := c.Prime(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Prime: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Prime of %d: %d\n", number, msg.Result)
	}
}
