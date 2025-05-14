package auth

import (
	"user-service/internal/domain/user"
)

// Configuration is an alias for a function that will take in a pointer to a Service and modify it
type Configuration func(s *Service) error

// Service is an implementation of the Service
type Service struct {
	userRepository user.Repository
	secretKey      []byte
}

// New takes a variable amount of Configuration functions and returns a new Service
// Each Configuration will be called in the order they are passed in
func New(configs ...Configuration) (s *Service, err error) {
	// Add the service
	s = &Service{}

	// Apply all Configurations passed in
	for _, cfg := range configs {
		// Pass the service into the configuration function
		if err = cfg(s); err != nil {
			return
		}
	}
	return
}

func WithUserRepository(userRepository user.Repository, secretKey []byte) Configuration {
	// Add the book repository, if we needed parameters, such as connection strings they could be inputted here
	return func(s *Service) error {
		s.userRepository = userRepository
		s.secretKey = secretKey
		return nil
	}
}
