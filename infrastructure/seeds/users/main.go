package main

import (
	"golang-ddd/domain/entity"
	"golang-ddd/infrastructure/eloquent"
	"golang-ddd/utils/helpers"
	"fmt"
)

func main() {
	r, err := eloquent.NewRepositories()

	r.Statement("TRUNCATE TABLE users;")

	if err != nil {
		helpers.HandleErrors(err)
	}

	count := 10

	r.User.Create(entity.User{
		Name:     "Admin",
		Email:    "admin@gmail.com",
		Password: "123456",
		UserType: "admin",
	})

	for i := 1; i <= count; i++ {
		r.User.Create(entity.User{
			Name:     fmt.Sprintf("User Number %d", i),
			Email:    fmt.Sprintf("user%d@test.com", i),
			Password: "123456",
		})
	}
}
