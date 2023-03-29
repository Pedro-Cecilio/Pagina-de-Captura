package main

import (
	"net/http"

	"github.com/Pedro-Cecilio/Pagina-de-captura/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
