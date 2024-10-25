package types

type CreateUserPayload struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}
type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Password  string `json:"_"`
	Role      string `json:"role,omitempty"`
}
