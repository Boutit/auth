package server

import (
	"context"
	"fmt"
	"time"

	"github.com/Boutit/auth/api"
	"github.com/Boutit/auth/internal/config"
	"github.com/golang-jwt/jwt"
)

func (s authServiceServer) CreateToken(ctx context.Context, req *api.CreateTokenRequest) (*api.CreateTokenResponse, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			// standardized claims
			"aud": "api",
			"nbf": now.Unix(),
			"iat": now.Unix(),
			"exp": now.Add(time.Minute).Unix(),
			"iss": config.GetAppConfig().Host,

			// user is custom claim for the validated user
			"userId": req.UserId,

			// roles is a list of roles attached to the user
			// it shows that claims can have more complex value types
			"roles": req.Roles,
	})

	tokenString, err := token.SignedString([]byte(config.GetTokenConfig().AccessTokenPrivateKey))

	if (err != nil){
		return &api.CreateTokenResponse{}, fmt.Errorf("unable to sign token for userId: %s due to error: %t", req.UserId, err)
	}

	return &api.CreateTokenResponse{
		Token: tokenString,
	}, nil
}

func (s authServiceServer) ValidateToken(ctx context.Context, req *api.ValidateTokenRequest) (*api.ValidateTokenResponse, error) {
	return &api.ValidateTokenResponse{}, nil
}

func (s authServiceServer) RefreshAccessToken(ctx context.Context, req *api.RefreshAccessTokenRequest) (*api.RefreshAccessTokenResponse, error) {
	return &api.RefreshAccessTokenResponse{}, nil
}