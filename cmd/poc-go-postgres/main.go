package main

import (
	"strings"

	"github.com/gs-open-provider/poc-go-postgres/internal/configs"
	"github.com/gs-open-provider/poc-go-postgres/internal/database"
	"github.com/gs-open-provider/poc-go-postgres/internal/logger"
	"github.com/gs-open-provider/poc-go-postgres/models"

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

	database.SelectAllUsers(db)
	database.SelectOneUser(db, 3)

	user := models.User{
		ID:     5,
		Name:   "5th User",
		Emails: []string{"email@gmail.com", "email2@gmail.com"},
	}
	err := database.AddNewUser(db, &user)
	if err != nil {
		logger.Log.Error(err.Error())
		if strings.Contains(err.Error(), "duplicate") {
			logger.Log.Info("User already exists, so skipping insert..")
		}
	}
	database.SelectOneUser(db, 5)

}
