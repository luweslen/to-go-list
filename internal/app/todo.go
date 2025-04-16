package app

type Todo struct {
	Id          int
	Name        string
	Description string
	Done        bool
}

func (todo *Todo) Toggle() *Todo {
	todo.Done = !todo.Done

	return todo
}

func (todo *Todo) SetName(name string) *Todo {
	todo.Name = name

	return todo
}

func (todo *Todo) SetDescription(description string) *Todo {
	todo.Description = description

	return todo
}

func FindById(todos []Todo, id int) *Todo {
	for i := 0; i < len(todos); i++ {
		if todos[i].Id == id {
			return &todos[i]
		}
	}
	return nil
}
