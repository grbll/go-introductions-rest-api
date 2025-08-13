package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/grbll/go-introductions-rest-api/service"
)

type LoginRequest struct {
	Email string `json:"email"`
}

type LoginResponse struct {
	Message string `json:"message"`
}

type AuthHandler struct {
	userService *service.UserService
}

func NewAuthHandler(us *service.UserService) *AuthHandler {
	return &AuthHandler{userService: us}
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

	exists, err := h.userService.IsUserRegistered(r.Context(), login.Email)
	if err != nil {
		log.Printf("failed to check user registration for %q: %v", login.Email, err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "User not Registered", http.StatusNotFound)
	}

	response.Message = fmt.Sprintf("Welcome %s", login.Email)

	json.NewEncoder(w).Encode(response)
	return
}
