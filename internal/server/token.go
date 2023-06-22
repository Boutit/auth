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
		return nil, fmt.Errorf("could not parse the token private key: %w", err)
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
		return nil, fmt.Errorf("could not create token: %w", err)
	}

	return &auth.CreateTokenResponse{
		Token: *td.Token,
		TokenUuid: td.TokenUUID,
		UserId: td.UserID,
		ExpiresIn: *td.ExpiresIn,
	}, nil
}

func (s authServiceServer) ValidateToken(ctx context.Context, req *auth.ValidateTokenRequest) (*auth.ValidateTokenResponse, error) {

	decodedPublicKey, err := base64.StdEncoding.DecodeString(config.GetTokenConfig().AccessTokenPublicKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode token public key: %w", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil{
		return nil, fmt.Errorf("could not parse the token public key: %w", err)
	}
	
	parsedToken, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if (err != nil){
		return nil, fmt.Errorf("the token could not be parsed: %w", err)
	}

	tokenClaims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	// check for an audience claim
	aud, ok := tokenClaims["aud"] 
	if !ok {
			return nil, fmt.Errorf("token has no audience claim")
	}

	// check that audience is from the issuer
	if aud != "api" {
			return nil, fmt.Errorf("token has the wrong audience claim")
	}

	if (err != nil){
		return nil, fmt.Errorf("unable to validate token due to error: %t", err)
	}

	return &auth.ValidateTokenResponse{
		TokenUuid: fmt.Sprint(tokenClaims["token_uuid"]),
		UserId: fmt.Sprint(tokenClaims["sub"]),
	}, nil
}

func (s authServiceServer) RefreshAccessToken(ctx context.Context, req *auth.RefreshAccessTokenRequest) (*auth.RefreshAccessTokenResponse, error) {
	return &auth.RefreshAccessTokenResponse{}, nil
}