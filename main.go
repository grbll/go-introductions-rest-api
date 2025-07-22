package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	// "time"

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

	fmt.Println("Connection to timestampdb succesufll!")

	_, err = db.Exec(`INSERT INTO users (user_email) VALUES (?)`, "newuser@example.com")
	if err != nil {
		log.Fatalf("Insert failed: %v", err)
	}

	http.HandleFunc("/login", handleLogin)

	log.Println("Goapp 0.0.1 Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}
	var email string = r.FormValue("email")

	fmt.Fprintf(w, "Welcome %v", email)
	return
}
