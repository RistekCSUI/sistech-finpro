package utils

import (
	"fmt"
	"github.com/RistekCSUI/sistech-finpro/shared/config"
	"github.com/golang-jwt/jwt"
	"time"
)

type JWT struct {
	Config *config.EnvConfig
}

func (j *JWT) GenerateToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 8).Unix(),
		Issuer:    userId,
	})

	return token.SignedString([]byte(j.Config.JWTSecret))
}

func (j *JWT) ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.Config.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (j *JWT) ExtractTokenData(tokenString string) string {
	token, _ := j.ParseToken(tokenString)
	claims, _ := token.Claims.(jwt.MapClaims)
	return claims["iss"].(string)
}

func NewJWT(config *config.EnvConfig) JWT {
	return JWT{
		Config: config,
	}
}
