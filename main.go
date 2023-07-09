package main;

import (
	"log"
	"net"

	"github.com/apinanyogaratnam/jwt-grpc-server/jwt"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("Failed to listen on port 9000", err);
	}

	grpcServer := grpc.NewServer()

	jwtServer := jwt.Server{}
	jwt.RegisterJWTServiceServer(grpcServer, &jwtServer)

	log.Println("Starting gRPC server on port 9000...");
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve gRPC server over port 9000");
	}
}
