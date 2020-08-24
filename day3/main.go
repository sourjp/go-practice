package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sourjp/go-practice/day3/application"
	"github.com/sourjp/go-practice/day3/handler"
	"github.com/sourjp/go-practice/day3/infra"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost user=root password=root dbname=users sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}

	ui := infra.NewUserInfra(db)
	ua := application.NewUserApplication(ui)
	uh := handler.NewUserHandler(ua)

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/users/:id", uh.Get)
		v1.POST("/users/", uh.Create)
		v1.PUT("/users/:id", uh.Update)
		v1.DELETE("/users/:id", uh.Delete)
	}
	r.Run(":8080")
}
