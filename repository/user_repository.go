package repository

import (
	"github.com/Nasa28/hotel-room-reservation/types"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(user types.CreateUserPayload) error
	GetUserByEmail(email string) (*types.User, error)
	GetUserByID(id int) (*types.User, error)
	UpdateUser(id int, payload types.User) error
	DeleteUser(id int) error
}

type UserStore struct {
	db *sqlx.DB
}

func NewUserStore(db *sqlx.DB) *UserStore {
	return &UserStore{db: db}
}

func (s *UserStore) CreateUser(user types.CreateUserPayload) error {
	return nil
}

func (s *UserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (s *UserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (s *UserStore) UpdateUser(id int, payload types.User) error {
	return nil
}

func (s *UserStore) DeleteUser(id int) error {
	return nil
}
