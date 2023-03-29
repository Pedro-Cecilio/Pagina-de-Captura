package controllers

import (
	"net/http"
	"text/template"

	"github.com/Pedro-Cecilio/Pagina-de-captura/models"
)

var temp = template.Must(template.ParseGlob("templates/*html"))

func Index(w http.ResponseWriter, r *http.Request){
	temp.ExecuteTemplate(w, "index.html", nil)
}

func Insert(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		nome := r.FormValue("nome")
		email := r.FormValue("email")
		estado := r.FormValue("estado")

		models.InserirNoDb(nome, email, estado)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}