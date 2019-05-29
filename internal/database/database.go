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
