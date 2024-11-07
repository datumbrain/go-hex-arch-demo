package api

import (
	"encoding/json"
	"hexashop/internal/adapter/http/rq"
	"hexashop/internal/port"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type User struct {
	userSvc port.UserSvc
}

func NewUser(userSvc port.UserSvc) *User {
	return &User{
		userSvc: userSvc,
	}
}

// CreateUser handler
func (api *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	var User rq.User
	if err := json.NewDecoder(r.Body).Decode(&User); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	req := rq.DomainUser(User)
	if err := api.userSvc.CreateUser(r.Context(), &req); err != nil {
		http.Error(w, "Failed to create User", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(User)
}

// GetUser handler
func (api *User) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)
	User, err := api.userSvc.GetUser(r.Context(), idInt)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(User)
}

// UpdateUser handler
func (api *User) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)

	var User rq.User
	if err := json.NewDecoder(r.Body).Decode(&User); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	req := rq.DomainUser(User)
	if err := api.userSvc.UpdateUser(r.Context(), idInt, &req); err != nil {
		http.Error(w, "Failed to update User", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(User)
}

// DeleteUser handler
func (api *User) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)
	if err := api.userSvc.DeleteUser(r.Context(), idInt); err != nil {
		http.Error(w, "Failed to delete User", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
