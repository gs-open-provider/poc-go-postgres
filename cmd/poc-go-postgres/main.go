package main

import (
	"github.com/gs-open-provider/poc-go-postgres/internal/configs"
	"github.com/gs-open-provider/poc-go-postgres/internal/logger"
)

func main() {
	// Initialize Viper across the application
	configs.InitializeViper()

	// Initialize Logger across the application
	logger.InitializeZapCustomLogger()
}
