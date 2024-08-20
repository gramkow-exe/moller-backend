package main

import (
	"fmt"

	"moller-backend/src/entities"
	"moller-backend/src/routes"
)

func main() {
	// Entities package
	user := entities.User{
		Id:   1,
		Nome: "ugo",
	}
	fmt.Println(user)

	// Routes package
	routes.Test()
}
