package database

import (
	"fmt"
	"golang-ddd/utils/helpers"

	"github.com/jinzhu/gorm"
)

// MysqlConnect connects to mysql database.
func MysqlConnect() *gorm.DB {
	db, err := gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			helpers.GetEnv("DB_USERNAME"),
			helpers.GetEnv("DB_PASSWORD"),
			helpers.GetEnv("DB_HOST"),
			helpers.GetEnv("DB_PORT"),
			helpers.GetEnv("DB_DATABASE"),
		),
	)

	helpers.HandleErrors(err)

	db.LogMode(helpers.GetEnv("ENABLE_DB_LOG_MODE") == "enable")

	return db
}
