package database

import (
	"ddd/utils/helpers"
	"fmt"

	"github.com/jinzhu/gorm"
)

// MysqlConnect connects to mysql database.
func MysqlConnect() *gorm.DB {
	db, err := gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
			helpers.GetEnv("DB_USERNAME"),
			helpers.GetEnv("DB_PASSWORD"),
			helpers.GetEnv("DB_DATABASE"),
		),
	)

	helpers.HandleErrors(err)

	db.LogMode(true)

	return db
}
