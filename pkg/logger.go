package pkg

import (
	"log/slog"
	"os"
)

func NewLogger() *slog.Logger {
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	logger := slog.New(jsonHandler)
	return logger
}
