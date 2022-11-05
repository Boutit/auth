package store

import (
	_ "database/sql"

	"github.com/Boutit/auth/internal/config"
	"github.com/jmoiron/sqlx"
)

type postgresStore struct {
	conn *sqlx.DB
}

type TokenStore interface {

}

func CreatePostgresStore(cfg config.Config) (TokenStore, error) {
	conn, err := sqlx.Connect("postgres", cfg.PostgresConfig.GetConnectionString())
	if err != nil {
		return nil, err
	}
	return &postgresStore{
		conn: conn,
	}, nil
}
