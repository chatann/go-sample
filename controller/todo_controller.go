package controller

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	"github.com/chatann/go-sample/controller/dto"
	"github.com/chatann/go-sample/model/entity"
	"github.com/chatann/go-sample/model/repository"
)

type TodoController interface {
	GetTodos(w http.ResponseWriter, r *http.Request)
	PostTodo(w http.ResponseWriter, r *http.Request)
	PutTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
}

type todoController struct {
	tr repository.TodoRepository
}

func NewTodoController(tr repository.TodoRepository) TodoController {
	return &todoController{tr}
}

func (tc *todoController) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := tc.tr.GetTodos()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var todoResponses []dto.TodoResponse
	for _, todo := range todos {
		todoResponses = append(todoResponses, dto.TodoResponse{
			Id: todo.Id,
			Title: todo.Title,
			Content: todo.Content,
		})
	}
	var todosResponse dto.TodosResponse
	todosResponse.Todos = todoResponses
	output, _ := json.Marshal(todosResponse.Todos)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (tc *todoController) PostTodo(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var todoRequest dto.TodoRequest
	err := json.Unmarshal(body, &todoRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	todo := entity.TodoEntity{
		Title: todoRequest.Title,
		Content: todoRequest.Content,
	}
	id, err := tc.tr.InsertTodo(todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", r.Host + r.URL.Path + strconv.Itoa(id))
	w.WriteHeader(http.StatusCreated)
}

func (tc *todoController) PutTodo(w http.ResponseWriter, r *http.Request) {
	todoId, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var todoRequest dto.TodoRequest
	err = json.Unmarshal(body, &todoRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	todo := entity.TodoEntity{
		Id: todoId,
		Title: todoRequest.Title,
		Content: todoRequest.Content,
	}
	err = tc.tr.UpdateTodo(todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (tc *todoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = tc.tr.DeleteTodo(todoId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}