package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/abhay786-20/fraud-auth-service/internal/config"
)

type Postgres struct {
	DB *sqlx.DB
}

func NewPostgres(cfg config.DatabaseConfig) (*Postgres, error) {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.MaxLifetime)

	return &Postgres{DB: db}, nil
}

func (p *Postgres) Ping() error {
	return p.DB.Ping()
}

func (p *Postgres) Close() error {
	return p.DB.Close()
}
