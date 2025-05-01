package sqlite

import (
	"database/sql"

	"github.com/luweslen/to-go-list/internal/domain/entities"
)

func todoFromDB(row *sql.Row) (entities.Todo, error) {
	var t entities.Todo

	err := row.Scan(&t.Id, &t.Name, &t.Description, &t.Done)

	if err != nil {
		return t, err
	}

	return t, nil
}
