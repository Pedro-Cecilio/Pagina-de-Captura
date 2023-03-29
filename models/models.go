package models

import (
	"log"

	"github.com/Pedro-Cecilio/Pagina-de-captura/db"
)

func InserirNoDb(nome, email, estado string) {
	db := db.ConnectDb()
	preparedRequisition, err := db.Prepare("INSERT INTO dados(nome, email, estado) VALUES (?, ?, ?)")
	if err != nil{
		log.Println("Erro ao inserir produto:", err)
	}
	preparedRequisition.Exec(nome, email, estado)
}