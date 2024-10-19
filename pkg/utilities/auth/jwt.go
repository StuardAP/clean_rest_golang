package auth

import (
	"time"

	"github.com/StuardAP/clean_rest_golang/pkg/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)


func CreateJWT(secret []byte, id uuid.UUID) (string, error) {
	expirationTime := time.Second * time.Duration(config.Envs.JWTConfig.JWTExpirationTime)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256 , jwt.MapClaims{
		"id": id.String(),
		"expiredAt": time.Now().Add(expirationTime).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}




