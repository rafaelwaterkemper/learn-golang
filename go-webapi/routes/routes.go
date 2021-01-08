package routes

import (
	"learn-golang/go-webapi/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
