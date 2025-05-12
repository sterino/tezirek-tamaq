package client

import (
	"api-gateway/proto/gen/userpb"
	"google.golang.org/grpc"
	"log"
)

type UserClient userpb.UserServiceClient

func InitUserClient(addr string) UserClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to user service: %v", err)
	}
	return userpb.NewUserServiceClient(conn)
}
