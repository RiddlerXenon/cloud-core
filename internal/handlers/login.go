package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RiddlerXenon/cloud-core/internal/auth"
	"go.uber.org/zap"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		zap.S().Errorw("Method not allowed", "method", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		zap.S().Errorw("Failed to decode request", "error", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	token, err := auth.Login(req.Username, req.Password)
	if err != nil {
		zap.S().Errorw("Login failed", "error", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	resp := loginResponse{
		Token: token,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		zap.S().Errorw("Failed to encode response", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	zap.S().Infow("User logged in", "username", req.Username)
	zap.S().Debugw("Login request", "request", req)
	zap.S().Debugw("Login response", "response", resp)
}
