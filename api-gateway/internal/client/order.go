package client

import (
	"api-gateway/proto/gen/orderpb"
	"google.golang.org/grpc"
	"log"
)

type OrderClient orderpb.OrderServiceClient

func InitOrderClient(addr string) OrderClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to order service: %v", err)
	}
	return orderpb.NewOrderServiceClient(conn)
}
