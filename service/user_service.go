package service

import (
	"encoding/json"
	"example/GoApp/model"
	"example/GoApp/repository"
	"log"
	"net/http"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("decoding of user failed :  %d\n", r.Body)

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repo.CreateUser(&user); err != nil {
		log.Printf("user creation failed for user email : %s  due to error  : %v", user.Email, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated) // write a status code for response header
	json.NewEncoder(w).Encode(user)   //  write a  response   , this writes to ResponseWriter hence no return statement required
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.repo.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users) //  write a  response   , this writes to ResponseWriter hence no return statement required
}
