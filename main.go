package main

import (
	"fmt"

	"github.com/luweslen/to-go-list/internal/app"
)

type Todo = app.Todo

func main() {
	todos := []Todo{}

	todo := Todo{
		Id:          1,
		Name:        "Estudar Go",
		Description: "Vamos estudar Go",
		Done:        false,
	}

	todos = append(todos, todo)
	oldTodo := app.FindById(todos, 1)
	oldTodo.Toggle().SetDescription("Estudei Go").SetName("Estudando Go")
	app.FindById(todos, 1)

	fmt.Println(&todo.Id)
	fmt.Println(&todos[0].Id)
	fmt.Println(&oldTodo.Id)
}
