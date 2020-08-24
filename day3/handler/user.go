package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sourjp/go-practice/day3/application"
	"github.com/sourjp/go-practice/day3/domain"
)

type userHandler struct {
	ua application.UserApplication
}

func NewUserHandler(ua application.UserApplication) userHandler {
	return userHandler{ua: ua}
}

// Get: GET /users/:id
func (uh userHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}

	usr, err := uh.ua.GetByID(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalServerError"})
		return
	}
	c.JSON(http.StatusOK, usr)
}

// Create: POST /users
func (uh userHandler) Create(c *gin.Context) {
	var u domain.User
	if err := c.ShouldBindJSON(&u); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}
	if err := uh.ua.Create(u); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalServerError"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "craeted"})
}

// Delete: DELETE /users/:id
func (uh userHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}

	if err := uh.ua.Delete(id); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalServerError"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

// Update: PUT /users/:id
func (uh userHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}

	var u domain.User
	if err := c.ShouldBindJSON(&u); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}
	log.Println(c)

	if err := uh.ua.ChangeProfile(id, u); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalServerError"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}
