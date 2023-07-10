package jwt

import (
	"log"
	"time"
	"fmt"
	"database/sql"

	jwt_protobuf "github.com/apinanyogaratnam/jwt-grpc-server/jwt-protobuf/jwt"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/net/context"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Server struct {}

func InsertEventIntoDatabase(userId int, event string) {
	db, err := sql.Open("pgx", "postgres://postgres:postgres@host.docker.internal:5432/postgres");
	if err != nil {
		log.Fatalln("Unable to connect to database: ", err);
	}

	if err := db.Ping(); err != nil {
		log.Fatalln("Unable to ping database: ", err);
	}

	row := db.QueryRow("INSERT INTO jwt_logs (user_id, event) VALUES ($1, $2)", userId, event);
	if err := row.Err(); err != nil {
		log.Fatalln("Unable to insert into database: ", err);
	}
}

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

func ValidateToken(signedToken string) (bool, int) {
	secretKey := "secret"
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return false, -1
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := int(claims["userId"].(float64))
		return true, userId
	}
	return false, -1
}

func (*Server) GetToken(ctx context.Context, message *jwt_protobuf.GetTokenRequest) (*jwt_protobuf.GetTokenResponse, error) {
	log.Println("Received message body from client: ", message.Id);
	userId := int(message.Id)

	InsertEventIntoDatabase(userId, "get-token");

	token, err := GenerateJWT(userId)
	if err != nil {
		return nil, err
	}
	return &jwt_protobuf.GetTokenResponse{Token: token}, nil;
}

func (*Server) ValidateToken(ctx context.Context, message *jwt_protobuf.ValidateTokenRequest) (*jwt_protobuf.ValidateTokenResponse, error) {
	log.Println("Received message body from client: ", message.Token);

	valid, userId := ValidateToken(message.Token)

	InsertEventIntoDatabase(userId, "validate-token")

	return &jwt_protobuf.ValidateTokenResponse{Valid: valid}, nil;
}
