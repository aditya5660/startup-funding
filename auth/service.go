package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
}
type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = []byte("S3MU4B154_K3Y")

// generate token
func (s *jwtService) GenerateToken(userID int) (string, error) {
	// create payload
	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// signature token
	signinToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signinToken, err
	}
	return signinToken, nil
}

// verify token
