package main

import (
	"go-backend-template/internal/config"
	"go-backend-template/internal/handler"
	"go-backend-template/internal/repository"
	"go-backend-template/internal/service"
	"go-backend-template/pkg/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.Load()

	db, err := database.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Инициализация репозиториев
	userRepo := repository.NewUserRepository(db)
	profileRepo := repository.NewProfileRepository(db)

	// Инициализация сервисов
	userService := service.NewUserService(userRepo)
	profileService := service.NewProfileService(profileRepo)

	// Инициализация обработчиков
	userHandler := handler.NewUserHandler(userService)
	profileHandler := handler.NewProfileHandler(profileService)

	// Настройка маршрутизатора
	r := mux.NewRouter()

	// Пользователи
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	// Профили
	r.HandleFunc("/profiles", profileHandler.CreateProfile).Methods("POST")
	r.HandleFunc("/profiles", profileHandler.GetAllProfiles).Methods("GET")
	r.HandleFunc("/profiles/{id}", profileHandler.GetProfile).Methods("GET")
	r.HandleFunc("/profiles/{id}", profileHandler.UpdateProfile).Methods("PUT")
	r.HandleFunc("/profiles/{id}", profileHandler.DeleteProfile).Methods("DELETE")

	log.Printf("Server is running on port %s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, r))
}
