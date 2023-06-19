package domain

type UserStore interface {
	Create(user User) error
	Delete(email string) error
	Update(user User) error
	GetOne(email string) (User, error)
	UpdateProminent(isProminent bool, host string) error
	GetAllProminent() ([]string, error)
}
