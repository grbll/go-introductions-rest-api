package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	// "log"
	"net/http"
)

type AuthHandler struct {
	DB *sql.DB
}
type LoginRequest struct {
	Email string `json:"email"`
}

type LoginResponse struct {
	Message string `json:"message"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}

	var login LoginRequest
	var response LoginResponse

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	response.Message = fmt.Sprintf("Welcome %s", login.Email)

	json.NewEncoder(w).Encode(response)
	return
}
