package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sourjp/go-practice/day2/models"
)

func Router() {
	r := gin.Default()

	db, err := models.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	v1 := r.Group("/api/v1/")
	{
		th := NewTODOHandler(models.NewTODOList(db))
		v1.GET("/todos", th.GetItems)
		v1.GET("/todos/:id", th.Get)
		v1.POST("/todos", th.Create)
		v1.PUT("/todos/:id", th.Update)
		v1.DELETE("/todos/:id", th.Delete)
	}
	r.Run(":8080")
}
