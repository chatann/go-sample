package main

import (
	"fmt"
	"net/http"

	"github.com/chatann/go-sample/controller"
	"github.com/chatann/go-sample/model/repository"
)

var tr = repository.NewTodoRepository()
var tc = controller.NewTodoController(tr)
var ro = controller.NewRouter(tc)

func main() {
	server := http.Server{Addr: ":8081"}
	http.HandleFunc("/todos/", ro.HandleTodoRequest)
	server.ListenAndServe()
	fmt.Println("Server running on port 8081")
}