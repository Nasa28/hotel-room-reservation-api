package handlers

import (
	"fmt"
	"net/http"

	"github.com/Nasa28/hotel-room-reservation/repository"
)

type UserHandler struct {
	UserRepository repository.UserRepository
}

func NewUserHandler(repository repository.UserRepository) *UserHandler {
	return &UserHandler{UserRepository: repository}
}
func (h *UserHandler) RegisterRoutes(router *http.ServeMux) {

	router.HandleFunc("GET /users", h.handleGetUser)
}

func (h *UserHandler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get users")
}
