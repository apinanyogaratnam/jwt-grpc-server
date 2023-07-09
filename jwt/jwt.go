package jwt;

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {}

func (*Server) GetToken(ctx context.Context, message *JWTRequest) (*JWTResponse, error) {
	log.Println("Received message body from client: ", message.Id);
	return &JWTResponse{Token: "Hello from the server!"}, nil;
}
