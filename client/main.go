package main

import (
	"context"
	"flag"
	"time"

	"github.com/dtorannpu/grpc-go-example/sample"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Err(err).Msg("did not connect")
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal().Err(err).Msg("could not close connection")
		}
	}(conn)
	c := sample.NewSampleServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Sample(ctx, &sample.SampleRequest{})
	if err != nil {
		log.Fatal().Err(err).Msg("could not sample request")
	}
	log.Printf("response %s", r.GetMessage())
}
