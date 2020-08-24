package infra_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/sourjp/go-practice/day3/domain"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sourjp/go-practice/day3/infra"
)

var db *sql.DB

func TestUserInfra_GetByID(t *testing.T) {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	// db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER  PRIMARY KEY AUTOINCREMENT NOT NULL,
		name VARCHAR(255),
		password VARCHAR(255)
	)`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`INSERT INTO users (name, password) VALUES ('Tom', 'Pass123')`)
	if err != nil {
		panic(err)
	}
	tests := []struct {
		name string
		id   int
		want domain.User
	}{
		{name: "Create Normal User", id: 1, want: domain.User{ID: 1, Name: "Tom", Password: "Pass123"}},
	}

	ui := infra.NewUserInfra(db)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usr, err := ui.GetByID(test.id)
			if err != nil {
				t.Errorf("got errors: %s", err)
			}
			if test.want.ID != usr.ID {
				t.Errorf("exepct: %d, but got: %d", test.want.ID, usr.ID)
			}
			if test.want.Name != usr.Name {
				t.Errorf("exepct: %s, but got: %s", test.want.Name, usr.Name)
			}
			if test.want.Password != usr.Password {
				t.Errorf("exepct: %s, but got: %s", test.want.Password, usr.Password)
			}
		})
	}
}
