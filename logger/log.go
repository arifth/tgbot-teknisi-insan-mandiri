package logger

import (
	"fmt"
	"go.uber.org/zap"
	"time"
)

func NewLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			fmt.Println(err)
		}
	}(logger) // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "google.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", "google.com")
	return logger
}
