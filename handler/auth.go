package handler

import (
	"encoding/json"
	"fmt"
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
		writeJSONError(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}

	var login LoginRequest
	var response LoginResponse

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		writeJSONError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	exists, err := h.userService.IsUserRegistered(r.Context(), login.Email)
	if err != nil {
		writeJSONError(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if !exists {
		writeJSONError(w, "User not Registered", http.StatusNotFound)
		return
	}

	response.Message = fmt.Sprintf("Welcome %s", login.Email)
	json.NewEncoder(w).Encode(response)
	return
}

func writeJSONError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var login LoginRequest
	var response LoginResponse

	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		writeJSONError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	exists, err := h.userService.IsUserRegistered(r.Context(), login.Email)
	if err != nil {
		writeJSONError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if exists {
		writeJSONError(w, "Userer already registered", http.StatusConflict)
		return
	}

	err = h.userService.RegisterUser(r.Context(), login.Email)
	if err != nil {
		writeJSONError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response.Message = fmt.Sprintf("%s successfully registered!", login.Email)
	json.NewEncoder(w).Encode(response)
	return
}
