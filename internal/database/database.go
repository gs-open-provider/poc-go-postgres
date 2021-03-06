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
			logger.Log.Error("Error:" + err.Error())
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

// SelectOneUser Function
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

// AddNewUser Function
func AddNewUser(db *pg.DB, user *models.User) error {
	if err := db.Insert(user); err != nil {
		logger.Log.Error("Insert Error: " + err.Error())
		return err
	}
	logger.Log.Info("User added..")
	return nil
}

// UpdateUser Function
func UpdateUser(db *pg.DB, modifiedUser *models.User) error {
	err := db.Update(modifiedUser)
	if err != nil {
		logger.Log.Error("Update error: " + err.Error())
		return err
	}
	logger.Log.Info("User updated..")
	return nil
}

// DeleteUser Function
func DeleteUser(db *pg.DB, id int64) error {
	user := &models.User{ID: id}
	err := db.Delete(user)
	if err != nil {
		logger.Log.Error("Delete error: " + err.Error())
		return err
	}
	logger.Log.Info("User deleted..")
	return nil
}
