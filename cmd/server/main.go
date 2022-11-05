package main

import (
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/lib/pq"

	"github.com/Boutit/auth/api"
	"github.com/Boutit/auth/internal/config"
	"github.com/Boutit/auth/internal/server"
	"github.com/Boutit/auth/internal/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	cfg := config.GetConfig(env)
	
	// use net package's Listen method to create a server
	h := cfg.AppConfig.Host
	appPort := cfg.AppConfig.Port
	appValues := []interface{}{h, appPort}
	appConnStr := fmt.Sprintf("%s:%d", appValues...)

	lis, err := net.Listen("tcp", appConnStr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	log.Printf("listening on %s", appConnStr)
	 
	// create a gRPC server
	g := grpc.NewServer()
	reflection.Register(g)

	// Create store
	tokenStore, err := store.CreatePostgresStore(cfg)
	if err != nil {
		log.Fatalf("failed to connect to user user store: %v", err)
	}
	
	// create a service server
	u := server.NewAuthServiceServer(tokenStore)
	
	// register the service server with the gRPC server
	api.RegisterAuthServiceServer(g, u)

	if err := g.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}