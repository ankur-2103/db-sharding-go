package handler

import (
	"db-sharding/internal/services"
	"db-sharding/internal/store"
	"db-sharding/internal/user"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		Service: userService,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newUser user.User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err = h.Service.CreateUser(newUser)

	if err != nil {
		switch {
		case errors.Is(err, store.ErrInvalidUserID):
			http.Error(w, err.Error(), http.StatusBadRequest)

		case errors.Is(err, store.ErrDuplicateUserID):
			http.Error(w, err.Error(), http.StatusConflict)

		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {

	userIDString := r.PathValue("userId")

	if userIDString == "" {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	userID, errInt := strconv.Atoi(userIDString)

	if errInt != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	var updateUser user.User

	err := json.NewDecoder(r.Body).Decode(&updateUser)

	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	res, err := h.Service.UpdateUser(userID, updateUser)

	if err != nil {
		switch {
		case errors.Is(err, store.ErrInvalidUserID):
			http.Error(w, err.Error(), http.StatusBadRequest)

		case errors.Is(err, store.ErrUserNotFound):
			http.Error(w, err.Error(), http.StatusNotFound)

		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		return
	}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userIDString := r.PathValue("userId")

	if userIDString == "" {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDString)

	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	res, err := h.Service.GetUser(userID)

	if err != nil {
		switch {
		case errors.Is(err, store.ErrInvalidUserID):
			http.Error(w, err.Error(), http.StatusBadRequest)

		case errors.Is(err, store.ErrUserNotFound):
			http.Error(w, err.Error(), http.StatusNotFound)

		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		return
	}
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userIDString := r.PathValue("userId")

	if userIDString == "" {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDString)

	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	res, err := h.Service.DeleteUser(userID)

	if err != nil {
		switch {
		case errors.Is(err, store.ErrInvalidUserID):
			http.Error(w, err.Error(), http.StatusBadRequest)

		case errors.Is(err, store.ErrUserNotFound):
			http.Error(w, err.Error(), http.StatusNotFound)

		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		return
	}
}