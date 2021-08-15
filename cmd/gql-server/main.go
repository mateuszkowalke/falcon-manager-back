package main

import (
	"log"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to initiate zap logger: %s\n", err)
	}
	logger.Info("Logger working correctly")
}
