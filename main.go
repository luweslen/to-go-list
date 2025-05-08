package main

import (
	"log"

	"github.com/luweslen/to-go-list/internal/domain/entities"
	"github.com/luweslen/to-go-list/internal/infraestructure/db/sqlite"
)

func main() {
	sqlite.OpenDB()

	todo := entities.Todo{
		Id:          1,
		Name:        "Estudar Go",
		Description: "Vamos estudar Go",
		Done:        false,
	}

	repository := sqlite.TodoRepository{}

	result := repository.Create(todo)

	result.Toggle()
	result.Name = "Alterado"

	repository.Update(result.Id, result)

	todo2, err := repository.GetAll()

	if err != nil {
		log.Print("Erro ao buscar todo:", err)
	}

	log.Print("Todos encontrado:", todo2)

}
