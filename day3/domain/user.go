package domain

// User is User model.
type User struct {
	ID       int    `db:"id", json:"id"`
	Name     string `db:"name", json:"name"`
	Password string `db:"password", json:"password"`
}

// UserRepository is repository for User model.
type UserRepository interface {
	GetByID(id int) (User, error)
	Create(u User) error
	Update(u User) error
	Delete(id int) error
}
