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
	log.Print("Teste", result)

	fetched, err := repository.GetById(9)
	if err != nil {
		log.Print("Erro ao buscar todo:", err)
	}
	log.Print("Todo encontrado:", fetched)

	allTodos, err := repository.GetAll()
	if err != nil {
		log.Print("Erro ao buscar todos:", err)
	}
	log.Print("Todos encontrados:", allTodos)

	var teste = fetched.Toggle()
	teste2 := repository.Create(*teste)
	log.Print("Teste", teste2)

}
