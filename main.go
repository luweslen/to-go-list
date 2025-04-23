package main

import (
	"database/sql"
	"log"

	"github.com/joho/godotenv"
	"github.com/luweslen/to-go-list/internal/app"
	_ "github.com/mattn/go-sqlite3"
)

type Todo = app.Todo

var DB *sql.DB

func loadEnv() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func initDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./app.db") // Open a connection to the SQLite database file named app.db
	if err != nil {
		log.Fatal(err) // Log an error and stop the program if the database can't be opened
	}

	// SQL statement to create the todos table if it doesn't exist
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS todos (
	 id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	 title TEXT,
	 description TEXT,
	 done BOOL
	);`

	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Error creating table: %q: %s\n", err, sqlStmt) // Log an error if table creation fails
	}
}

func main() {
	loadEnv()
	initDB()

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

	log.Print(&todo.Id)
	log.Print(&todos[0].Id)
	log.Print(&oldTodo.Id)
}
