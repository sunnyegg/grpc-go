package main

import (
	"io"
	"log"

	pb "github.com/sunnyegg/grpc-go/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg function was invoked")

	res := float64(0)
	total := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Result: res / float64(total),
			})
		}

		if err != nil {
			log.Fatalf("error while reading the client stream: %v\n", err)
		}

		log.Printf("Receiving: %v\n", req)
		res += float64(req.Number)
		total += 1
	}
}
