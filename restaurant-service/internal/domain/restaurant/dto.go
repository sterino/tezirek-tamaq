package restaurant

type Request struct {
	Name     string   `json:"name" validate:"required"`
	Address  string   `json:"address" validate:"required"`
	Phone    string   `json:"phone" validate:"required"`
	OrderIDs []string `json:"order_ids,omitempty"`
}

type Response struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	Phone    string   `json:"phone"`
	OrderIDs []string `json:"order_ids,omitempty"`
}

func ParseFromEntity(e Entity) Response {
	return Response{
		ID:       e.ID,
		Name:     e.Name,
		Address:  e.Address,
		Phone:    e.Phone,
		OrderIDs: e.OrderIDs,
	}
}

func ParseFromEntities(data []Entity) []Response {
	result := make([]Response, len(data))
	for i, r := range data {
		result[i] = ParseFromEntity(r)
	}
	return result
}
