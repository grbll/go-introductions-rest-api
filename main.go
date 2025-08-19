package main

import (
	"log"
	"net/http"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/grbll/go-introductions-rest-api/handler"

	"github.com/grbll/go-introductions-rest-api/service/user"

	"github.com/grbll/go-introductions-rest-api/repository/user/mysql"
)

var version string = "0.0.64"

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

	var userRepository = mysqluserrepo.NewMySQLUserRepository(db)
	var userService = userservice.NewUserService(userRepository)
	var authHandler *handler.AuthHandler = handler.NewAuthHandler(userService)

	http.HandleFunc("/login", authHandler.Login)
	http.HandleFunc("/register", authHandler.Register)

	log.Printf("Goapp %v Listening on http://localhost:8080", version)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
