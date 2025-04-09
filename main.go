package main

import "fmt"

type Todo struct {
	id          int
	name        string
	description string
	done        bool
}

func (todo *Todo) toggle() *Todo {
	todo.done = !todo.done

	return todo
}

func (todo *Todo) setName(name string) *Todo {
	todo.name = name

	return todo
}

func (todo *Todo) setDescription(description string) *Todo {
	todo.description = description

	return todo
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
	oldTodo := findById(todos, 1)
	oldTodo.toggle().setDescription("Estudei Go").setName("Estudando Go")
	findById(todos, 1)

	fmt.Println(&todo.id)
	fmt.Println(&todos[0].id)
	fmt.Println(&oldTodo.id)
}
