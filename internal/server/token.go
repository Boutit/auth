package server

import (
	"context"

	"github.com/Boutit/auth/api"
)

func (s authServiceServer) CreateToken(ctx context.Context, req *api.CreateTokenRequest) (*api.CreateTokenResponse, error) {
	return &api.CreateTokenResponse{}, nil
}

func (s authServiceServer) ValidateToken(ctx context.Context, req *api.ValidateTokenRequest) (*api.ValidateTokenResponse, error) {
	return &api.ValidateTokenResponse{}, nil
}

func (s authServiceServer) RefreshAccessToken(ctx context.Context, req *api.RefreshAccessTokenRequest) (*api.RefreshAccessTokenResponse, error) {
	return &api.RefreshAccessTokenResponse{}, nil
}