package database

import (
	"context"
	"database/sql"
	"fmt"
	"go-backend-template/internal/config"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func NewPostgresConnection(cfg *config.Config) (*sql.DB, error) {
	// Временная жесткая настройка (подставьте свои значения)
	//надо будет попробовать давать значения не в явном виде, но пока что у меня не получается сделать это рабочим
	connStr := "host=localhost port=5432 user=postgres password=chara dbname=go_backend sslmode=disable"

	log.Println("Используется строка подключения:", connStr) // Убедитесь, что пароль подставился верно

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть подключение: %v", err)
	}

	// Проверка подключения с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ошибка ping: %v", err)
	}

	return db, nil
}
