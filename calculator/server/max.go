package main

import (
	"io"
	"log"

	pb "github.com/sunnyegg/grpc-go/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max function was invoked")

	var max int64 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		if req.Number > max {
			max = req.Number
		}

		log.Printf("Receiving: %v\n", req.Number)

		err = stream.Send(&pb.MaxResponse{
			Result: max,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}
