package infra

import (
	"database/sql"

	"github.com/sourjp/go-practice/day3/domain"
)

type userInfra struct {
	db *sql.DB
}

func NewUserInfra(db *sql.DB) domain.UserRepository {
	return &userInfra{db: db}
}

func (ui *userInfra) GetByID(id int) (domain.User, error) {
	const sql = "SELECT * FROM users WHERE id = $1"

	var u domain.User
	if err := ui.db.QueryRow(sql, id).Scan(&u.ID, &u.Name, &u.Password); err != nil {
		return domain.User{}, err
	}
	return u, nil
}

func (ui *userInfra) Create(u domain.User) error {
	const sql = "INSERT INTO users (name, password) VALUES($1, $2)"

	if _, err := ui.db.Exec(sql, u.Name, u.Password); err != nil {
		return err
	}
	return nil
}

func (ui *userInfra) Update(u domain.User) error {
	const sql = "UPDATE users SET name = $1, password = $2 WHERE id = $3"
	if _, err := ui.db.Exec(sql, u.Name, u.Password, u.ID); err != nil {
		return err
	}
	return nil
}

func (ui *userInfra) Delete(id int) error {
	const sql = "DELETE FROM users WHERE id = $1"
	if _, err := ui.db.Exec(sql, id); err != nil {
		return err
	}
	return nil
}
