package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type AuthHandler struct {
	DB *sql.DB
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("hallo")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}
	var data struct {
		Email string `json:"email"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	var response struct {
		Message string `json:"message"`
	}

	response.Message = fmt.Sprintf("Welcome %s", data.Email)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return
}
