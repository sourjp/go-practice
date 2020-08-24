package application_test

import (
	"errors"
	"testing"

	"github.com/sourjp/go-practice/day3/application"

	"github.com/sourjp/go-practice/day3/domain"
)

func TestUserApplication_GetByID(t *testing.T) {
	tests := []struct {
		name string
		id   int
		want application.UserDTO
	}{
		{name: "Get Normal User", id: 1, want: application.UserDTO{ID: 1, Name: "Tom"}},
	}

	ua := application.NewUserApplication(TestUserRepository{})
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usr, err := ua.GetByID(test.id)
			if err != nil {
				t.Fatal(err)
			}
			if usr.ID != test.want.ID {
				t.Errorf("exepct: %d, but got: %d", test.want.ID, usr.ID)
			}
			if usr.Name != test.want.Name {
				t.Errorf("exepct: %s, but got: %s", test.want.Name, usr.Name)
			}
		})
	}
}

type TestUserRepository struct{}

func (tur TestUserRepository) GetByID(id int) (domain.User, error) {
	if id != 1 {
		return domain.User{}, errors.New("There is no id")
	}
	u := domain.User{
		ID:       1,
		Name:     "Tom",
		Password: "test123",
	}
	return u, nil
}

func (tur TestUserRepository) Create(u domain.User) error {
	return nil
}

func (tur TestUserRepository) Update(u domain.User) error {
	return nil
}

func (tur TestUserRepository) Delete(id int) error {
	return nil
}
