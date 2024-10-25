package api

import (
	"net/http"

	"github.com/Nasa28/hotel-room-reservation/handlers"
	"github.com/Nasa28/hotel-room-reservation/repository"
	"github.com/jmoiron/sqlx"
)

type APIServer struct {
	addr string
	db   *sqlx.DB
}

func NewAPIServer(addr string, db *sqlx.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	v1 := http.NewServeMux()

	// We will register routes below
	userStore := repository.NewUserStore(s.db)
	userHandler := handlers.NewUserHandler(userStore)
	userHandler.RegisterRoutes(v1)
	router.Handle("/api/v1/", http.StripPrefix("/api/v1", v1))
	return http.ListenAndServe(s.addr, router)
}
