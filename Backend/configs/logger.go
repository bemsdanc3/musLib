package configs

import (
	"log/slog"
	"os"
)

func InitLogger() *slog.Logger {
	logFile, err := os.OpenFile("configs/logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	logger := slog.New(slog.NewJSONHandler(logFile))
	return logger
}
