package main

import (
	"log"

	pb "github.com/sunnyegg/grpc-go/calculator/proto"
)

func (s *Server) Prime(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeServer) error {
	log.Printf("Prime was invoked with: %v\n", in)

	for i := int32(1); i < in.Number; i++ {
		if in.Number%i == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: i,
			})
		}
	}

	return nil
}
