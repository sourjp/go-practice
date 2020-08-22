package models

import (
	"database/sql"
	"fmt"
	"time"
)

type TODOList struct {
	db *sql.DB
}

func NewTODOList(db *sql.DB) *TODOList {
	return &TODOList{db: db}
}

type TODO struct {
	ID         int        `db:"id" json:"id"`
	Title      string     `db:"title" json:"title"`
	Message    string     `db:"message" json:"message"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	FinishedAt *time.Time `db:"finished_at" json:"finished_at,omitempty"`
}

func (tl *TODOList) Create(t TODO) error {
	if t.CreatedAt.IsZero() {
		t.CreatedAt = time.Now()
	}
	fmt.Println(t)
	const sql = "INSERT INTO todo (title, message, created_at, finished_at) VALUES ($1, $2, $3, $4)"
	_, err := tl.db.Exec(sql, t.Title, t.Message, t.CreatedAt, t.FinishedAt)
	if err != nil {
		return err
	}

	return nil
}

func (tl *TODOList) Get(id int) (TODO, error) {
	var t TODO
	const sql = "SELECT * FROM todo WHERE id = $1"
	err := tl.db.QueryRow(sql, id).Scan(&t.ID, &t.Title, &t.Message, &t.CreatedAt, &t.UpdatedAt, &t.FinishedAt)
	if err != nil {
		return TODO{}, err
	}
	return t, nil
}

func (tl *TODOList) GetItems(limit int) ([]TODO, error) {
	const sql = "SELECT * FROM todo ORDER BY id DESC LIMIT $1"
	rows, err := tl.db.Query(sql, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []TODO
	for rows.Next() {
		var t TODO
		rows.Scan(&t.ID, &t.Title, &t.Message, &t.CreatedAt, &t.UpdatedAt, &t.FinishedAt)
		todos = append(todos, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}

func (tl *TODOList) Update(t TODO, id int) error {
	const sql = "UPDATE todo SET title = $1, message = $2, updated_at = $3 WHERE id = $4"
	updatedAt := time.Now()
	t.UpdatedAt = &updatedAt
	_, err := tl.db.Exec(sql, t.Title, t.Message, t.UpdatedAt, id)
	if err != nil {
		return err
	}
	return err
}

func (tl *TODOList) Delete(id int) error {
	const sql = "DELETE FROM todo WHERE id = $1"
	_, err := tl.db.Exec(sql, id)
	if err != nil {
		return err
	}
	return nil
}
