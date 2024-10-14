package app

import (
	"Backend/internal/delivery/http"
	"Backend/internal/repository"
	"Backend/internal/usecases"
	"Backend/pkg/db"
	"github.com/gin-gonic/gin"
	"log"
)

func Run() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	userRepo := repository.NewUserRepository(database)
	userUsecase := usecases.NewUseCase(userRepo)

	r := gin.Default()

	userHandler := http.NewUserHandler(userUsecase)
	r.GET("/users", userHandler.GetAllUsers)
	r.POST("/users", userHandler.CreateUser)

	if err := r.Run(":5002"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
