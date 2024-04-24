package main

import (
	"fmt"
	papi "github.com/qtpeters/pendiente/api"
)

func main() {
	fmt.Println("Test")
	api := papi.NewRestApi()
	api.EnableAddTodo()
	api.Run()
}
