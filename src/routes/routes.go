package routes

import (
	"log"
	"moller-backend/src/controllers"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
