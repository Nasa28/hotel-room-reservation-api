package handlers

import (
	"fmt"
	"net/http"

	"github.com/Nasa28/hotel-room-reservation/auth"
	"github.com/Nasa28/hotel-room-reservation/config"
	"github.com/Nasa28/hotel-room-reservation/repository"
	"github.com/Nasa28/hotel-room-reservation/types"
	"github.com/Nasa28/hotel-room-reservation/utils"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	UserRepository repository.UserRepository
}

func NewUserHandler(repository repository.UserRepository) *UserHandler {
	return &UserHandler{UserRepository: repository}
}
func (h *UserHandler) RegisterRoutes(router *http.ServeMux) {

	router.HandleFunc("GET /users", h.handleGetUser)
	router.HandleFunc("POST /register", h.handleCreateUser)
}

func (h *UserHandler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get users")
}

func (h *UserHandler) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	result, err := h.UserRepository.GetUserByEmail(payload.Email)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error retrieving user: %v", err))
		return
	}

	if (result != types.User{}) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}
	hashedPassword, err := auth.HashedPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.UserRepository.CreateUser(types.CreateUserPayload{
		Email:        payload.Email,
		FirstName:    payload.FirstName,
		LastName:     payload.LastName,
		Password:     hashedPassword,
		PhoneNumber:  payload.PhoneNumber,
		StreetName:   payload.StreetName,
		StreetNumber: payload.StreetNumber,
		City:         payload.City,
		State:        payload.State,
		Country:      payload.Country,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	u, err := h.UserRepository.GetUserByEmail(payload.Email)
	if err != nil {

		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	secret := []byte(config.Env.JWTSecret)
	token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, map[string]string{"token": token})
}
