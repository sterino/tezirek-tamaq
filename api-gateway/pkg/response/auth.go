package response

import "api-gateway/proto/gen/authpb"

type AuthResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

func FromAuthProto(a *authpb.TokenResponse) AuthResponse {
	return AuthResponse{
		Token: a.GetToken(),
	}
}
