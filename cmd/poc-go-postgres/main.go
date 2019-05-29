package main

import (
	"os"
	"strconv"
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

	err := database.CreateSchema(db)
	if err != nil {
		if strings.Contains(err.Error(), "database \""+dbName+"\" does not exist") {
			logger.Log.Info("Database " + dbName + " does not exist..")
			os.Exit(1)
		}
	}

	database.SelectAllUsers(db)
	database.SelectOneUser(db, 3)

	user := models.User{
		ID:     5,
		Name:   "4th User",
		Emails: []string{"email@gmail.com", "email2@gmail.com"},
	}
	err = database.AddNewUser(db, &user)
	if err != nil {
		logger.Log.Error(err.Error())
		if strings.Contains(err.Error(), "duplicate") {
			logger.Log.Info("User already exists, so skipping insert..")
		}
	}
	database.SelectOneUser(db, 5)

	const modifyingUserID int64 = 5
	modifiedUser := models.User{
		ID:     modifyingUserID,
		Name:   "5th User",
		Emails: []string{"email@gmail.com", "email2@gmail.com"},
	}
	err = database.UpdateUser(db, &modifiedUser)
	if err != nil && strings.Contains(err.Error(), "no rows in result set") {
		logger.Log.Warn("User with the same Primary Key does not exist..")
	} else {
		database.SelectOneUser(db, modifyingUserID)
	}

	const deletingUserID int64 = 5
	err = database.DeleteUser(db, deletingUserID)
	if err != nil && strings.Contains(err.Error(), "no rows in result set") {
		logger.Log.Warn("User with the same Primary Key does not exist..")
	}
	err = database.SelectOneUser(db, deletingUserID)
	if err != nil && strings.Contains(err.Error(), "no rows in result set") {
		logger.Log.Warn("User with ID of " + strconv.FormatInt(deletingUserID, 10) + " does not exist..")
	}

	db.Close()

}
