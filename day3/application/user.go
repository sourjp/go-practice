package application

import (
	"log"

	"github.com/sourjp/go-practice/day3/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserApplication interface {
	GetByID(id int) (UserDTO, error)
	Create(usr domain.User) error
	ChangeProfile(id int, usr domain.User) error
	Delete(id int) error
}

type userApplication struct {
	ur domain.UserRepository
}

type UserDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewUserApplication(ur domain.UserRepository) UserApplication {
	return &userApplication{ur: ur}
}

func (ua *userApplication) GetByID(id int) (UserDTO, error) {
	usr, err := ua.ur.GetByID(id)
	if err != nil {
		return UserDTO{}, err
	}
	return UserDTO{ID: usr.ID, Name: usr.Name}, nil
}

func (ua *userApplication) Create(usr domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), 12)
	if err != nil {
		return err
	}
	usr.Password = string(hash)
	if err := ua.ur.Create(usr); err != nil {
		return err
	}
	return nil
}

func (ua *userApplication) ChangeProfile(id int, usr domain.User) error {
	getUsr, err := ua.ur.GetByID(id)
	if err != nil {
		return err
	}

	if len(usr.Name) != 0 {
		getUsr.Name = usr.Name
	}
	if len(usr.Password) != 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), 12)
		if err != nil {
			return err
		}
		getUsr.Password = string(hash)
	}

	log.Println(id, usr)
	if err := ua.ur.Update(getUsr); err != nil {
		return err
	}
	return nil
}

func (ua *userApplication) Delete(id int) error {
	if err := ua.ur.Delete(id); err != nil {
		return err
	}
	return nil
}
