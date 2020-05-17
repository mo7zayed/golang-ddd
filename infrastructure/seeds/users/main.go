package main

import (
	"ddd/domain/entity"
	"ddd/infrastructure/persistence"
	"ddd/utils/helpers"
	"fmt"
)

func main() {
	r, err := persistence.NewRepositories()

	r.Statement("TRUNCATE TABLE users;")

	if err != nil {
		helpers.HandleErrors(err)
	}

	count := 20

	r.User.Create(entity.User{
		Name:     "Admin",
		Email:    "admin@gmail.com",
		Password: "123456",
	})

	for i := 1; i <= count; i++ {
		r.User.Create(entity.User{
			Name:     fmt.Sprintf("User Number %d", i),
			Email:    fmt.Sprintf("user%d@test.com", i),
			Password: "123456",
		})
	}
}
