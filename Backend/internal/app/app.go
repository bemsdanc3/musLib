package app

import (
	"Backend/configs"
	db2 "Backend/pkg/db"
)

func StartServer() {
	logger := configs.InitLogger()

	database, err := db2.InitDb()
	if err != nil {
		logger.Error("failed to connect to database:", err)
	}
	defer database.Close()

}
