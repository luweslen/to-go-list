package entities

type Todo struct {
	Id          int64
	Name        string
	Description string
	Done        bool
}

func (t *Todo) Toggle() *Todo {
	t.Done = !t.Done

	return t
}

func (t *Todo) SetName(name string) *Todo {
	t.Name = name

	return t
}

func (t *Todo) SetDescription(description string) *Todo {
	t.Description = description

	return t
}
