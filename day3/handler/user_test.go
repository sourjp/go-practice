package handler_test

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sourjp/go-practice/day3/domain"
	"github.com/sourjp/go-practice/day3/handler"

	"github.com/sourjp/go-practice/day3/application"
)

func TestUserHandler_Get(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		wantCode int
	}{
		{name: "Get Expect 200", id: "1", wantCode: 200},
		{name: "Get Expect 500", id: "2", wantCode: 500},
	}
	gin.SetMode(gin.ReleaseMode)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			req := httptest.NewRequest("GET", "/api/v1/users/1", nil)
			c.Request = req

			var tua TestUserApplication
			uh := handler.NewUserHandler(tua)

			p := gin.Param{Key: "id", Value: test.id}
			c.Params = []gin.Param{p}
			uh.Get(c)

			if test.wantCode != w.Code {
				t.Errorf("exepct: %d, but got: %d", test.wantCode, w.Code)
			}
		})
	}
}

type TestUserApplication struct{}

func (ua TestUserApplication) GetByID(id int) (application.UserDTO, error) {
	if id == 1 {
		usr := application.UserDTO{ID: 1, Name: "Tom"}
		return usr, nil
	}

	return application.UserDTO{}, errors.New("test")
}

func (ua TestUserApplication) Create(usr domain.User) error {
	return nil
}

func (ua TestUserApplication) ChangeProfile(id int, usr domain.User) error {
	return nil
}

func (ua TestUserApplication) Delete(id int) error {
	return nil
}
