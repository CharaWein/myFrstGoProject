package main

import (
	"log"
	"net/http"
	"time"

	"go-backend-template/internal/config"
	"go-backend-template/internal/handler"
	"go-backend-template/internal/repository"
	"go-backend-template/internal/service"
	"go-backend-template/pkg/database"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database connection
	db, err := database.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	profileRepo := repository.NewProfileRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo)
	profileService := service.NewProfileService(profileRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)
	profileHandler := handler.NewProfileHandler(profileService)

	// Create router
	r := mux.NewRouter()

	// User routes
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	// Profile routes
	r.HandleFunc("/profiles", profileHandler.CreateProfile).Methods("POST")
	r.HandleFunc("/profiles", profileHandler.GetAllProfiles).Methods("GET")
	r.HandleFunc("/profiles/{id}", profileHandler.GetProfile).Methods("GET")
	r.HandleFunc("/profiles/{id}", profileHandler.UpdateProfile).Methods("PUT")
	r.HandleFunc("/profiles/{id}", profileHandler.DeleteProfile).Methods("DELETE")
	r.HandleFunc("/profiles/search", profileHandler.SearchProfile).Methods("GET")

	// Configure CORS with correct types
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           86400, // 24 hours in seconds
	})

	// Start server
	server := &http.Server{
		Handler:      corsHandler.Handler(r),
		Addr:         ":" + cfg.ServerPort,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server is running on port %s", cfg.ServerPort)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
