package main

import (
	"fmt"
	papi "github.com/qtpeters/pendiente/api"
	"github.com/qtpeters/pendiente/core/services"
)

func main() {
	fmt.Println("Test")

	todoManger := papi.TodoManager(services.NewTodoService())

	api := papi.NewRestApi(todoManger)
	api.EnableCreateTodo()
	api.EnableUpdateTodo()
	api.EnableDeleteTodo()
	api.Run()
}
