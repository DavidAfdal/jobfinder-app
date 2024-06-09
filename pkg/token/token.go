package token

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenUseCase interface {
	GenerateAccessToken(claims JwtCustomClaims) (string, error)
}
type tokenUseCase struct {
	secretKey string
}

type JwtCustomClaims struct {
	ID     uuid.UUID `json:"id"`
	Email  string `json:"email"`
	Address string `json:"address"`
	Role     string `json:"role"`
	PhoneNumber   string `json:"phone_number"`
	jwt.RegisteredClaims
}


func NewTokenUseCase(secretKey string) TokenUseCase {
	return &tokenUseCase{secretKey: secretKey}
}


func (t *tokenUseCase) GenerateAccessToken(claims JwtCustomClaims) (string,error) {
	plainToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	encodedToken, err := plainToken.SignedString([]byte(t.secretKey))

	if err != nil {
		return "", err
	}

	return encodedToken, nil
}
