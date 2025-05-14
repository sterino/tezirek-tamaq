package auth

import (
	"context"
	"errors"
	"log"
	"user-service/internal/domain/auth"
	"user-service/internal/domain/user"
	"user-service/pkg/jwt"
	"user-service/pkg/password"
)

// Register — регистрация нового пользователя
func (s *Service) Register(ctx context.Context, req user.Request) (string, error) {
	hashed, err := password.Generate(req.Password)
	if err != nil {
		return "", err
	}
	log.Printf("Password: %s", req.Password)
	userEntity := user.Entity{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashed,
		Role:     req.Role, // или по req.Role, если есть
	}

	return s.userRepository.Create(ctx, userEntity)
}

// Login — проверка данных и выдача JWT
func (s *Service) Login(ctx context.Context, req auth.Request) (string, int64, error) {
	u, err := s.userRepository.GetByEmail(ctx, req.Email)
	if err != nil {
		return "", 0, err
	}

	hashed, err := password.Generate(req.Password)

	if res := password.Compare(u.Password, hashed); res == false {
		log.Printf("Password comparison failed for user: %s", req.Email)
		log.Printf("Password comparison failed: %s", u.Password)
		log.Printf("Password comparison failed: %s", hashed)
		return "", 0, errors.New("invalid password")
	}
	token, expiredAt, err := jwt.Encode(jwt.JWT{
		u.ID,
		u.Email,
	}, s.secretKey)
	if err != nil {
		return "", 0, err
	}
	return *token, *expiredAt, nil
}
