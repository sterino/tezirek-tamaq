package response

import "api-gateway/proto/gen/userpb"

type UserResponse struct {
	ID    string `json:"id" example:"user-123"`
	Name  string `json:"name" example:"John Doe"`
	Email string `json:"email" example:"john@example.com"`
	Role  string `json:"role" example:"customer"`
}

func FromUserProto(u *userpb.UserResponse) UserResponse {
	return UserResponse{
		ID:    u.GetId(),
		Name:  u.GetName(),
		Email: u.GetEmail(),
		Role:  u.GetRole(),
	}
}

func FromUserProtoList(users []*userpb.UserResponse) []UserResponse {
	res := make([]UserResponse, len(users))
	for i, u := range users {
		res[i] = FromUserProto(u)
	}
	return res
}
