package client

import (
	"api-gateway/proto/gen/restaurantpb"
	"google.golang.org/grpc"
	"log"
)

type RestaurantClient restaurantpb.RestaurantServiceClient

func InitRestaurantClient(addr string) RestaurantClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to restaurant service: %v", err)
	}
	return restaurantpb.NewRestaurantServiceClient(conn)
}
