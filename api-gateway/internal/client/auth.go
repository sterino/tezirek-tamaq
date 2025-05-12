package client

import (
	"api-gateway/proto/gen/authpb"
	"google.golang.org/grpc"
	"log"
)

type AuthClient authpb.AuthServiceClient

func InitAuthClient(addr string) AuthClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to auth service: %v", err)
	}
	return authpb.NewAuthServiceClient(conn)
}
