package server

import (
	"context"
	"fmt"
	"time"

	"github.com/Boutit/auth/api/protos/boutit/auth"
	"github.com/Boutit/auth/internal/config"
	"github.com/golang-jwt/jwt"
)

func (s authServiceServer) CreateToken(ctx context.Context, req *auth.CreateTokenRequest) (*auth.CreateTokenResponse, error) {
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
		return &auth.CreateTokenResponse{}, fmt.Errorf("unable to sign token for userId: %s due to error: %t", req.UserId, err)
	}

	return &auth.CreateTokenResponse{
		Token: tokenString,
	}, nil
}

func (s authServiceServer) ValidateToken(ctx context.Context, req *auth.ValidateTokenRequest) (*auth.ValidateTokenResponse, error) {
	
	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetTokenConfig().AccessTokenPrivateKey), nil
	})

	if (err != nil){
		return nil, err
	}

	// check for an audience claim
	aud, ok := token.Claims.(jwt.MapClaims)["aud"]
	if !ok {
			return nil, fmt.Errorf("token had no audience claim")
	}

	// check that audience is from the issuer
	if aud != "api" {
			return nil, fmt.Errorf("token had the wrong audience claim")
	}

	tokenString, err := token.SignedString([]byte(config.GetTokenConfig().AccessTokenPublicKey))

	if (err != nil){
		return nil, fmt.Errorf("unable to validate token due to error: %t", err)
	}

	return &auth.ValidateTokenResponse{
		Token: tokenString,
	}, nil
}

func (s authServiceServer) RefreshAccessToken(ctx context.Context, req *auth.RefreshAccessTokenRequest) (*auth.RefreshAccessTokenResponse, error) {
	return &auth.RefreshAccessTokenResponse{}, nil
}