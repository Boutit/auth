package server

import (
	"github.com/Boutit/auth/api"
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