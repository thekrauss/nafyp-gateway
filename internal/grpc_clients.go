package internal

import (
	"log"
	"time"

	"github.com/thekrauss/nafyp-gateway/config"
	pbauth "github.com/thekrauss/nafyp-protos/gen/go/auth"
	"google.golang.org/grpc"
)

type GRPCClients struct {
	Auth pbauth.AuthServiceClient
}

func InitGRPCClients() *GRPCClients {
	addr := config.AppConfig.GRPCClients.AuthServiceAddr

	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("[gRPC] Connexion échouée vers AuthService (%s) : %v", addr, err)
	}

	log.Printf("[gRPC] Connecté à AuthService (%s) avec succès", addr)

	return &GRPCClients{
		Auth: pbauth.NewAuthServiceClient(conn),
	}
}
