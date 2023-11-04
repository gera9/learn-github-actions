package main

import (
	"gera9/learn-github-actions/pkg"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("message",
		// Structured context as strongly typed Field values.
		zap.String("urlpkg.GetGreeting():", pkg.GetGreeting()),
	)
}
