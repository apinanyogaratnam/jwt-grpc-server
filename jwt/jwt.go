package jwt

import (
	"log"
	"time"

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

func (*Server) GetToken(ctx context.Context, message *JWTRequest) (*JWTResponse, error) {
	log.Println("Received message body from client: ", message.Id);
	token, err := GenerateJWT(int(message.Id))
	if err != nil {
		return nil, err
	}
	return &JWTResponse{Token: token}, nil;
}
