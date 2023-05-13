package model

type User struct {
	Name     string  `json:"name"`
	Surname  string  `json:"surname"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Address  Address `json:"address"`
	Role     string  `json:"role"`
}
