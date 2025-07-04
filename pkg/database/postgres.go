package database

import (
	"database/sql"
	"go-backend-template/internal/config"

	_ "github.com/lib/pq"
)

func NewPostgresConnection(cfg *config.Config) (*sql.DB, error) {
	//connStr := fmt.Sprintf(
	//	"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	//		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	//)
	connStr := "host=localhost port=5432 user=postgres password=chara dbname=go_backend sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
