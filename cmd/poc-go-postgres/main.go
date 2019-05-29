package main

import (
	"github.com/gs-open-provider/poc-go-postgres/internal/configs"
	"github.com/gs-open-provider/poc-go-postgres/internal/database"
	"github.com/gs-open-provider/poc-go-postgres/internal/logger"

	"github.com/go-pg/pg"
	"github.com/spf13/viper"
)

func main() {
	// Initialize Viper across the application
	configs.InitializeViper()

	// Initialize Logger across the application
	logger.InitializeZapCustomLogger()

	dbusername := viper.GetString("pgdb.username")
	dbpassword := viper.GetString("pgdb.password")
	dbName := viper.GetString("pgdb.database")
	logger.Log.Info(dbName)

	db := pg.Connect(&pg.Options{
		User:     dbusername,
		Password: dbpassword,
		Database: dbName,
	})
	defer db.Close()

	database.CreateSchema(db)
}
