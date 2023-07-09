package jwt

import (
	"log"
	"time"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/net/context"
)

type Server struct {}

func GenerateJWT(userId int) (string, error) {
	secretKey := "secret"

	tokenContent := jwt.MapClaims{
		"userId": userId,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // The token will expire after 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenContent)
	signedToken, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateToken(signedToken string) (bool, string) {
	secretKey := "your_secret_key"
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return false, ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["userId"].(string)
		return true, userId
	}
	return false, ""
}

func (*Server) GetToken(ctx context.Context, message *JWTRequest) (*JWTResponse, error) {
	log.Println("Received message body from client: ", message.Id);
	token, err := GenerateJWT(int(message.Id))
	if err != nil {
		return nil, err
	}
	return &JWTResponse{Token: token}, nil;
}
