package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/sunnyegg/grpc-go/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.MaxRequest{
		{Number: 20},
		{Number: 1},
		{Number: 19},
		{Number: 293},
		{Number: 22},
		{Number: 220},
		{Number: 500},
		{Number: 5},
		{Number: 2},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
