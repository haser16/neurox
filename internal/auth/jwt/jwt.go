package auth_jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
	secret []byte
	ttl    time.Duration
}

func NewAuth(config Config) *Service {
	return &Service{
		secret: []byte(config.Secret),
		ttl:    config.Ttl,
	}
}

func (s *Service) Generate(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(s.ttl).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}
