package main

import (
	"database/sql"
	"golangfiles/api"

	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:dhanekula@tcp(localhost:3306)/Library"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	api.RegisterRouter(db)

	log.Println("serever strt on port 8080:")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
