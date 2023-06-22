package config

import (
	"time"

	"github.com/spf13/viper"
)

const (
	accessTokenPublicKey string = "token.access_token_public_key"
	accessTokenPrivateKey string = "token.access_token_private_key"
	accessTokenTtl string = "token.access_token_ttl"
)

type TokenConfig struct {
	AccessTokenPublicKey string
	AccessTokenPrivateKey string
	AccessTokenTtl time.Duration
}

func GetTokenConfig () TokenConfig {
	return TokenConfig{
		AccessTokenPublicKey: viper.GetString(accessTokenPublicKey),
		AccessTokenPrivateKey: viper.GetString(accessTokenPrivateKey),
		AccessTokenTtl: viper.GetDuration(accessTokenTtl),
	}
}