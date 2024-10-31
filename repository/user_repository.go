package repository

import (
	"database/sql"

	"github.com/Nasa28/hotel-room-reservation/types"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(user types.CreateUserPayload) error
	GetUserByEmail(email string) (types.User, error)
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
	query := `
		INSERT INTO users 
		(email, firstName, lastName, password, phoneNumber, streetName, streetNumber, city, state, country) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err := s.db.Exec(query,
		user.Email, user.FirstName, user.LastName, user.Password,
		user.PhoneNumber, user.StreetName, user.StreetNumber,
		user.City, user.State, user.Country,
	)

	if err != nil {
		return err
	}
	return nil
}

func (s *UserStore) GetUserByEmail(email string) (types.User, error) {
	query := `
		SELECT id, email FROM users WHERE email = $1
	`
	var user types.User
	err := s.db.Get(&user, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.User{}, nil
		}
		return types.User{}, err
	}

	return user, nil
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
