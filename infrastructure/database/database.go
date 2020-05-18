package database

import (
	"golang-ddd/utils/helpers"
	"errors"

	"github.com/jinzhu/gorm"

	//  importing db mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Connect connects to database based on spescfied database connection.
func Connect() *gorm.DB {
	if helpers.GetEnv("DB_CONNECTION") == "mysql" {
		return MysqlConnect()
	}

	helpers.HandleErrors(errors.New("there is no database connection spescfied in .env"))

	return &gorm.DB{}
}
