package config

import (
	_ "database/sql"
	"fmt"

	"github.com/spf13/viper"
)

const (
	postgresHostKey 		string = "db.host"
	postgresPortKey 		string = "db.port"
	postgresUserKey 		string = "db.user"
	postgresPasswordKey string = "db.password"
	postgresDBNameKey 	string = "db.dbname"
	postgresSSLMode			string = "db.sslmode"
)

type PostgresConfig struct {
	Host     	string
	Port     	string
	User		 	string
	Password 	string
	DBName   	string
	SSLMode		string
}

func GetPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host: 		viper.GetString(postgresHostKey),
		Port: 		viper.GetString(postgresPortKey),
		User:			viper.GetString(postgresUserKey),
		Password: viper.GetString(postgresPasswordKey),
		DBName: 	viper.GetString(postgresDBNameKey),
		SSLMode: 	viper.GetString(postgresSSLMode),
	}
}

func (p PostgresConfig) GetConnectionString() string {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", p.Host, p.Port, p.User, p.Password, p.DBName, p.SSLMode)
	return connStr
}