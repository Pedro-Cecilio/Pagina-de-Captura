package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb() *sql.DB {
	conect := "root:root@/pagina_de_captura"
	db, err := sql.Open("mysql", conect)
	if err != nil {
		log.Println("Erro ao acessar db:", err)
	}
	return db
}