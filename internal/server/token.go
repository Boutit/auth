package server

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/Boutit/auth/api/protos/boutit/auth"
	"github.com/Boutit/auth/internal/config"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type TokenDetails struct {
	Token     *string
	TokenUUID string
	UserID    string
	ExpiresIn *int64
}

func (s authServiceServer) CreateToken(ctx context.Context, req *auth.CreateTokenRequest) (*auth.CreateTokenResponse, error) {
	now := time.Now().UTC()
	td := &TokenDetails{
		ExpiresIn: new(int64),
		Token: new(string),
	}
	*td.ExpiresIn = now.Add(config.GetTokenConfig().AccessTokenTtl).Unix()
	td.TokenUUID = uuid.New().String()
	td.UserID = req.GetUserId()

	decodedPrivateKey, err := base64.StdEncoding.DecodeString(config.GetTokenConfig().AccessTokenPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode private key: %w", err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("there was a problem parsing the token private key: %w", err)
	}

	tokenClaims := make(jwt.MapClaims)
	tokenClaims["aud"] = "api"
	tokenClaims["nbf"] = now.Unix()
	tokenClaims["iat"] = now.Unix()
	tokenClaims["exp"] = td.ExpiresIn
	tokenClaims["iss"] = config.GetAppConfig().Host
	tokenClaims["token_uuid"] = td.TokenUUID
	tokenClaims["sub"] = req.GetUserId()
	*td.Token, err = jwt.NewWithClaims(jwt.SigningMethodRS256, tokenClaims).SignedString(key)
	if(err != nil) {
		return nil, fmt.Errorf("token could not be created due to error: %w", err)
	}

	return &auth.CreateTokenResponse{
		Token: *td.Token,
		TokenUuid: td.TokenUUID,
		UserId: td.UserID,
		ExpiresIn: *td.ExpiresIn,
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