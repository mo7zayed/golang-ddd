package main

import (
	"ddd/infrastructure/eloquent"
	"ddd/utils/helpers"
)

func main() {
	db, err := eloquent.NewRepositories()

	if err != nil {
		helpers.HandleErrors(err)
	}

	db.AutoMigrate()
}
