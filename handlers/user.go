package handlers

import (
	"fmt"
	"net/http"
)

type Handler struct{}

func (h *Handler) NewHandler() *Handler {
	return &Handler{}
}
func (h *Handler) RegisterRoutes(router *http.ServeMux) {

	router.HandleFunc("GET /users", h.handleGetUser)
}

func (h *Handler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get users")
}
