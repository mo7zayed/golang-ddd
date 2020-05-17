package main

import (
	"ddd/infrastructure/persistence"
	"ddd/utils/helpers"
)

func main() {
	db, err := persistence.NewRepositories()

	if err != nil {
		helpers.HandleErrors(err)
	}

	db.AutoMigrate()
}
