package routes

import (
	"log"
	"moller-backend/src/controllers"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/fmt", controllers.FmtWriter)
	http.HandleFunc("/json-direct", controllers.JsonDirectlyWrited)
	http.HandleFunc("/json-encoder", controllers.JsonEncode)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
