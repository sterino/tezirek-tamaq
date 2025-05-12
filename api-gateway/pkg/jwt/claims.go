package jwt

import "github.com/dgrijalva/jwt-go"

type Profile struct {
	Email string `json:"email"`
}

type Claims struct {
	Profile Profile
	jwt.StandardClaims
}
