package main

import (
	"golang-ddd/infrastructure/eloquent"
	"golang-ddd/utils/helpers"
)

func main() {
	db, err := eloquent.NewRepositories()

	if err != nil {
		helpers.HandleErrors(err)
	}

	db.AutoMigrate()
}
