package routes

import (
	"net/http"

	"github.com/Pedro-Cecilio/Pagina-de-captura/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/insert", controllers.Insert)
}