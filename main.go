package main

import "fmt"

type Todo struct {
	id          int
	name        string
	description string
	done        bool
}

func findById(todos []Todo, id int) *Todo {
	for i := 0; i < len(todos); i++ {
		if todos[i].id == id {
			return &todos[i]
		}
	}
	return nil
}

func main() {
	todos := []Todo{}

	todo := Todo{
		id:          1,
		name:        "Estudar Go",
		description: "Vamos estudar Go",
		done:        false,
	}

	todos = append(todos, todo)

	fmt.Println(findById(todos, 1))
}
