package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/sunnyegg/grpc-go/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet function was invoked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("error while reading the client stream: %v\n", err)
		}

		log.Printf("Receiving: %v\n", req)
		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}
