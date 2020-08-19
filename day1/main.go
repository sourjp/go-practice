package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/getall", HandleGetAll)
		v1.POST("/create", HandlePut)
	}
	log.Println(r.Run())
}

func HandleGetAll(c *gin.Context) {
	db, err := NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Can't connect to DB"})
		return
	}
	todos, err := GetAll(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err})
		return
	}

	if len(todos) == 0 {
		c.JSON(http.StatusOK, gin.H{"status": "there is no tods"})
		return
	}
	// c.JSON(http.StatusOK, todos)
	c.IndentedJSON(http.StatusOK, todos)
}

func HandlePut(c *gin.Context) {
	var t TODO
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}
	db, err := NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err})
		return
	}

	if err := Put(db, t); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "your recored recived"})
}

var (
	driver = "postgres"
	params = "host=localhost dbname=todo user=root password=root sslmode=disable"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open(driver, params)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

type TODO struct {
	ID         int        `db:"id" json:"id"`
	Title      string     `db:"title" json:"title"`
	Message    string     `db:"message" json:"message"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	FinishedAt *time.Time `db:"finished_at" json:"finished_at,omitempty"`
}

func Put(db *sql.DB, todo TODO) error {
	if todo.CreatedAt.IsZero() {
		todo.CreatedAt = time.Now()
	}

	r, err := db.Exec("INSERT INTO todo (title, message, created_at, finished_at) VALUES ($1, $2, $3, $4)", todo.Title, todo.Message, todo.CreatedAt, todo.FinishedAt)

	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}

func GetAll(db *sql.DB) ([]TODO, error) {
	rows, err := db.Query("SELECT * FROM todo")
	if err != nil {
		return nil, err
	}

	var t TODO
	var todos []TODO
	for rows.Next() {
		if err := rows.Scan(&t.ID, &t.Title, &t.Message, &t.CreatedAt, &t.FinishedAt); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}
