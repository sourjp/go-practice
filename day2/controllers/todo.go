package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sourjp/go-practice/day2/models"
)

func resWithErr(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func resNoErr(c *gin.Context, code int, message interface{}) {
	c.JSON(code, gin.H{"status": message})
}

type TODOHandler struct {
	tl *models.TODOList
}

func NewTODOHandler(tl *models.TODOList) *TODOHandler {
	return &TODOHandler{tl: tl}
}

func (th *TODOHandler) GetItems(c *gin.Context) {
	ls, ok := c.GetQuery("limit")
	if !ok {
		ls = "10"
	}
	li, err := strconv.Atoi(ls)
	if err != nil {
		log.Println(err)
		resWithErr(c, http.StatusBadRequest, "failed to parse your query")
		return
	}
	todos, err := th.tl.GetItems(li)
	if err != nil {
		log.Println(err)
		resWithErr(c, http.StatusInternalServerError, "failed to get items")
		return
	}
	if len(todos) == 0 {
		resNoErr(c, http.StatusOK, "there is no todos. create your todo!")
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (th *TODOHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		resWithErr(c, http.StatusBadRequest, "failed to parse URI")
		return
	}

	todo, err := th.tl.Get(id)
	if err != nil {
		log.Println(err)
		resWithErr(c, http.StatusInternalServerError, "failed to get item")
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (th *TODOHandler) Create(c *gin.Context) {
	var t models.TODO
	if err := c.ShouldBindJSON(&t); err != nil {
		log.Println(err)
		resWithErr(c, http.StatusBadRequest, "failed to parse your request")
		return
	}

	if err := th.tl.Create(t); err != nil {
		log.Println(err)
		resWithErr(c, http.StatusInternalServerError, "failed to create item")
		return
	}
	resNoErr(c, http.StatusOK, "created")
}

func (th *TODOHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		resWithErr(c, http.StatusBadRequest, "failed to parse URI")
		return
	}

	var t models.TODO
	if err := c.ShouldBindJSON(&t); err != nil {
		log.Println(err)
		resWithErr(c, http.StatusBadRequest, "failed to parse your request")
		return
	}

	if err := th.tl.Update(t, id); err != nil {
		log.Println(err)
		resWithErr(c, http.StatusInternalServerError, "failed to update item")
		return
	}

	resNoErr(c, http.StatusOK, "updated")
}

func (th *TODOHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		resWithErr(c, http.StatusBadRequest, "failed to parse URI")
		return
	}

	if err := th.tl.Delete(id); err != nil {
		log.Println(err)
		resWithErr(c, http.StatusInternalServerError, "failed to delete item")
	}

	resNoErr(c, http.StatusOK, "deleted")
}
