package main

import (
	"context"
	"log"

	pb "github.com/sunnyegg/grpc-go/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Printf("doSum was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstInt:  3,
		SecondInt: 10,
	})
	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Sum: %v\n", res.Result)
}
