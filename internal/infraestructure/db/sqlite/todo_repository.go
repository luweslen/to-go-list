package sqlite

import (
	"database/sql"
	"log"

	"github.com/luweslen/to-go-list/internal/domain/entities"
)

type TodoRepository struct{}

func (TodoRepository) Create(todo entities.Todo) entities.Todo {
	sqlSave := `INSERT INTO todos (title, description, done) VALUES (?, ?, ?)`

	result, err := DB.Exec(sqlSave, todo.Name, todo.Description, todo.Done)

	if err != nil {
		log.Print(err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		log.Print(err)
	}

	todo.Id = id

	return todo
}

func (TodoRepository) Delete(todoId int64) sql.Result {
	sqlDelete := `DELETE FROM todos WHERE id = ?`

	result, err := DB.Exec(sqlDelete, todoId)

	if err != nil {
		log.Print(err)
	}

	return result
}

func (TodoRepository) Update(id int64, todo entities.Todo) sql.Result {
	sqlUpdate := `UPDATE todos SET title = ?, description = ?, done = ? WHERE id = ?`

	result, err := DB.Exec(sqlUpdate, todo.Name, todo.Description, todo.Done, todo.Id)

	if err != nil {
		log.Print(err)
	}

	return result
}

func (TodoRepository) GetById(todoId int64) (entities.Todo, error) {
	sqlGetById := `SELECT id, title, description, done FROM todos WHERE id = ?`

	row := DB.QueryRow(sqlGetById, todoId)

	return todoFromDB(row)
}

func (TodoRepository) GetAll() ([]entities.Todo, error) {
	sqlGetAll := `SELECT id, title, description, done FROM todos`

	rows, err := DB.Query(sqlGetAll)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todos, err := todosFromDB(rows)

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}
