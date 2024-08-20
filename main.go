package main

import (
	"fmt"
	"moller-backend/src/routes"
)

func main() {
	fmt.Println("Iniciano o servidor Rest com Go!")
	routes.HandleRequest()
}
