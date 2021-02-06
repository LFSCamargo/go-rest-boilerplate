package auth

import (
	"errors"

	userModel "github.com/LFSCamargo/go-rest-boilerplate/database/models/user"
	"github.com/LFSCamargo/go-rest-boilerplate/env"
	"github.com/dgrijalva/jwt-go"
)

// GetUserFromToken - Gets the user from the token JWT
func GetUserFromToken(tokenStr string) (*userModel.User, error) {
	config := env.GetConfig()

	byteSecret := []byte(config.Jwt)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return byteSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["user_id"].(string)
		return userModel.FindByEmail(email)
	} else {
		return nil, errors.New("Invalid Token")
	}
}

// SignToken - Creates a new token with the user email
func SignToken(email string) (string, error) {
	config := env.GetConfig()
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = email
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	return at.SignedString([]byte(config.Jwt))
}
