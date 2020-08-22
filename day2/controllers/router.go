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

	v1 := r.Group("/api/v1/todo")
	{
		th := NewTODOHandler(models.NewTODOList(db))
		v1.GET("/getitems", th.GetItems)
		v1.GET("/get/:id", th.Get)
		v1.POST("/create", th.Create)
		v1.PUT("/update/:id", th.Update)
		v1.DELETE("/delete/:id", th.Delete)
	}
	r.Run(":8080")
}
