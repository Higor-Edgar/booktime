package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Higor-Edgar/booktime.git/internal/contract"
	"github.com/Higor-Edgar/booktime.git/internal/service"
)

type AuthHandler interface {
	SignUp(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	userService service.UserService
}

func (a *authHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var newUserDTO contract.NewUserDTO
	if err := json.NewDecoder(r.Body).Decode(&newUserDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := a.userService.CreateUser(&newUserDTO); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func NewAuthHandler(userService service.UserService) AuthHandler {
	return &authHandler{
		userService: userService,
	}
}
