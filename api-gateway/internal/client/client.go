package client

type Clients struct {
	Auth       AuthClient
	User       UserClient
	Order      OrderClient
	Restaurant RestaurantClient
}

func InitClients() *Clients {
	return &Clients{
		Auth:       InitAuthClient("user-service:50051"), // Auth реализован в user-service
		User:       InitUserClient("user-service:50051"),
		Order:      InitOrderClient("order-service:50052"),
		Restaurant: InitRestaurantClient("restaurant-service:50053"),
	}
}
