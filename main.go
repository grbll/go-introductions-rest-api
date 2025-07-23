package main

import (
	"database/sql"
	"github.com/grbll/go-introductions-rest-api/handler"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var dsn string = "goapp:goapp@tcp(mysql:3306)/timestampdb"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("DB open error: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("DB connect error: %v", err)
	}
	log.Println("Connection to timestampdb succesufll!")

	log.Println("Goapp 0.0.1 Listening on http://localhost:8080")

	var authHandler = &handler.AuthHandler{DB: db}

	http.HandleFunc("/login", authHandler.Login)
	log.Println(http.ListenAndServe(":8080", nil))
}
