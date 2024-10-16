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
	//инициализация репозиториев
	userRepo := repository.NewUserRepository(database)
	tracksRepo := repository.NewTrackRepository(database)
	//инициализация юзкейсов
	userUsecase := usecases.NewUseCase(userRepo)
	trackUsecase := usecases.NewTrackUseccase(tracksRepo)

	r := gin.Default()

	//обработка пользовательских путей
	userHandler := http.NewUserHandler(userUsecase)
	r.GET("/users", userHandler.GetAllUsers)
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users/:id", userHandler.GetUserByID)

	//инициализация путей треков
	trackHandler := http.NewTrackHandler(trackUsecase)
	r.GET("/users/:id/tracks", trackHandler.GetTracksByUserID)
	r.GET("/tracks", trackHandler.GetAllTracks)
	r.POST("/tracks", trackHandler.CreateTrack)

	if err := r.Run(":5002"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
