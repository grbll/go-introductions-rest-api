package main

import (
	"fmt"
	"log"

	"net/http"
)

func main() {
	http.HandleFunc("/login", handleLogin)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}
	var name string = r.FormValue("name")

	fmt.Fprintf(w, "Welcome %v", name)
	return
}
