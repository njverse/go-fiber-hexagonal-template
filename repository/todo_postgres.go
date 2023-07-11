package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type todoRepositoryPostgres struct {
	db *sqlx.DB
}

func NewTodoRepositoryPostgres(db *sqlx.DB) TodoRepository {
	return todoRepositoryPostgres{db: db}
}

func (r todoRepositoryPostgres) Create(todo Todo) (*Todo, error) {
	stmt, _ := r.db.PrepareNamed("INSERT INTO todos (type, name, status) VALUES (:type, :name, :status) RETURNING id")
	if err := stmt.Get(&todo, todo); err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r todoRepositoryPostgres) GetAll(todo Todo) ([]Todo, error) {
	query := "SELECT id, type, name, status FROM todos"

	whereClause := []string{}
	if todo.Name != "" {
		todo.Name = fmt.Sprintf("%%%s%%", todo.Name)
		whereClause = append(whereClause, "name like :name")
	}
	if todo.Type != "" {
		whereClause = append(whereClause, "type = :type")
	}
	if todo.Status != 0 {
		whereClause = append(whereClause, "status = :status")
	}
	if len(whereClause) > 0 {
		query += " WHERE " + strings.Join(whereClause, " and ")
	}

	var todos []Todo
	stmt, _ := r.db.PrepareNamed(query)
	if err := stmt.Select(&todos, todo); err != nil {
		return nil, err
	}

	return todos, nil
}
