package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/dtorannpu/grpc-go-example/sample"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	sample.UnimplementedSampleServiceServer
}

func (s *server) Sample(ctx context.Context, request *sample.SampleRequest) (*sample.SampleResponse, error) {
	return &sample.SampleResponse{Message: "Sample"}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}
	s := grpc.NewServer()
	sample.RegisterSampleServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("failed to server")
	}
}
