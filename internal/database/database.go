package database

import (
	"strings"

	"github.com/gs-open-provider/poc-go-postgres/internal/logger"
	"github.com/gs-open-provider/poc-go-postgres/models"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// CreateSchema Function
func CreateSchema(db *pg.DB) error {
	tables := []interface{}{(*models.User)(nil)}
	for _, model := range tables {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				logger.Log.Warn("Warning:" + err.Error())
			} else {
				return err
			}
		}
	}
	return nil
}

// SelectAllUsers Function
func SelectAllUsers(db *pg.DB) error {
	var users []models.User
	err := db.Model().Table("users").Select(&users)
	if err != nil {
		logger.Log.Error("Select error: " + err.Error())
		return err
	}

	for _, u := range users {
		logger.Log.Info(u.String())
	}
	return nil
}

// SelectOneUser exported
func SelectOneUser(db *pg.DB, id int64) error {
	user := &models.User{ID: id}
	err := db.Select(user)
	if err != nil {
		logger.Log.Error("Select error: " + err.Error())
		return err
	}

	logger.Log.Info(user.String())
	return nil
}
