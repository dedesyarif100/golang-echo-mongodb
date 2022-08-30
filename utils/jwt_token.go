package utils

import (
	"os"
	"time"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JwtTokenClaims struct {
	ID   string
	Name string
	jwt.StandardClaims
}

func GenerateAccessToken(id primitive.ObjectID, secret, email string) (string, error) {
	secret1 := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["email"] = email
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret1))
}