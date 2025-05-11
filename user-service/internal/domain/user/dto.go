package user

type Request struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role,omitempty"`
	Password string `json:"password,omitempty"`
}

type Response struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

// ParseFromEntity maps Entity to Response
func ParseFromEntity(e Entity) Response {
	return Response{
		ID:    e.ID,
		Name:  e.Name,
		Email: e.Email,
		Role:  e.Role,
	}
}

// ParseFromEntities maps []Entity to []Response
func ParseFromEntities(data []Entity) []Response {
	result := make([]Response, len(data))
	for i, u := range data {
		result[i] = ParseFromEntity(u)
	}
	return result
}

// NewFromUpdateRequest maps UpdateRequest to Entity
