package client

type Clients struct {
	Auth       AuthClient
	User       UserClient
	Order      OrderClient
	Restaurant RestaurantClient
}

func InitClients() *Clients {
	return &Clients{
		Auth:       InitAuthClient("localhost:50051"),
		User:       InitUserClient("localhost:50052"),
		Order:      InitOrderClient("localhost:50053"),
		Restaurant: InitRestaurantClient("localhost:50054"),
	}
}
