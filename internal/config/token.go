package config

import "github.com/spf13/viper"

const (
	accessTokenPublicKey string = "token.access_token_public_key"
	accessTokenPrivateKey string = "token.access_token_private_key"
)

type TokenConfig struct {
	AccessTokenPublicKey string
	AccessTokenPrivateKey string
}

func GetTokenConfig () TokenConfig {
	return TokenConfig{
		AccessTokenPublicKey: viper.GetString(accessTokenPublicKey),
		AccessTokenPrivateKey: viper.GetString(accessTokenPrivateKey),
	}
}