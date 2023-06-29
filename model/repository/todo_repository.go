package repository

import (
	"log"

	"github.com/chatann/go-sample/model/entity"
)

type TodoRepository interface {
	GetTodos() (todos []entity.TodoEntity, err error)
	InsertTodo(todo entity.TodoEntity) (id int, err error)
	UpdateTodo(todo entity.TodoEntity) (err error)
	DeleteTodo(id int) (err error)
}

type todoRepository struct {
}

func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

func (tr *todoRepository) GetTodos() (todos []entity.TodoEntity, err error) {
	todos = []entity.TodoEntity{}
	rows, err := Db.Query("SELECT id, title, content FROM todos ORDER BY id DESC")
	if err != nil {
		log.Print(err)
		return
	}
	for rows.Next() {
		var todo entity.TodoEntity
		err = rows.Scan(&todo.Id, &todo.Title, &todo.Content)
		if err != nil {
			log.Print(err)
			return
		}
		todos = append(todos, todo)
	}
	return
}

func (tr *todoRepository) InsertTodo(todo entity.TodoEntity) (id int, err error) {
	result, err := Db.Exec("INSERT INTO todos (title, content) VALUES (?, ?)", todo.Title, todo.Content)
	if err != nil {
		log.Print(err)
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		log.Print(err)
		return
	}
	id = int(id64)
	return
}

func (tr *todoRepository) UpdateTodo(todo entity.TodoEntity) (err error) {
	_, err = Db.Exec("UPDATE todos SET title = ?, content = ? WHERE id = ?", todo.Title, todo.Content, todo.Id)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (tr *todoRepository) DeleteTodo(id int) (err error) {
	_, err = Db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		log.Print(err)
		return
	}
	return
}