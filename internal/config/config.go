package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	AppConfig 			AppConfig
	PostgresConfig	PostgresConfig
	TokenConfig 		TokenConfig
}

func GetConfig(env string) Config {
	n := fmt.Sprintf("config.%s", env)

	viper.SetConfigName(n)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	
	err := viper.ReadInConfig()
	
	if err != nil {
		fmt.Println("Could not read config")
		fmt.Println(err)
	}

	return Config{
		AppConfig: GetAppConfig(),
		PostgresConfig: GetPostgresConfig(),
		TokenConfig: GetTokenConfig(),
	}
}