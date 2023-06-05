package server

import (
	api "github.com/Boutit/auth/api/protos/boutit/auth"
	"github.com/Boutit/auth/internal/store"
)

type AuthServiceServer interface {
	api.AuthServiceServer
}

type authServiceServer struct {
	api.UnimplementedAuthServiceServer
	tokenStore store.TokenStore
}

func NewAuthServiceServer(store store.TokenStore) AuthServiceServer {
	return &authServiceServer{
		tokenStore: store,
	}
}