package main

import (
	"context"
	"log"
	"time"

	pb "github.com/sunnyegg/grpc-go/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")

	req := []*pb.AverageRequest{
		{Number: 20},
		{Number: 40},
		{Number: 60},
	}

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Avg: %v\n", err)
	}

	for _, r := range req {
		log.Printf("Sending req: %v\n", r)
		stream.Send(r)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from Avg: %v\n", err)
	}

	log.Printf("Avg: %v\n", res.Result)
}
