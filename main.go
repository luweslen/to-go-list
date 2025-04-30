package main

import (
	"database/sql"
	"log"

	"github.com/joho/godotenv"
	"github.com/luweslen/to-go-list/internal/app"

	// _ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
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
	DB, err = sql.Open("sqlite", "./app.db") // Open a connection to the SQLite database file named app.db
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

func SaveTodo(todo Todo, DB *sql.DB) Todo {
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

func todoMapper(row *sql.Row) (Todo, error) {

	var t Todo
	err := row.Scan(&t.Id, &t.Name, &t.Description, &t.Done)
	if err != nil {
		return t, err
	}
	return t, nil

}

func GetTodoById(todoId int64, DB *sql.DB) (Todo, error) {
	sqlGetById := `SELECT id, title, description, done FROM todos WHERE id = ?`
	row := DB.QueryRow(sqlGetById, todoId)
	return todoMapper(row)
}

func GetTodos(DB *sql.DB) ([]Todo, error) {
	sqlGetAll := `SELECT id, title, description, done FROM todos`
	rows, err := DB.Query(sqlGetAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.Id, &t.Name, &t.Description, &t.Done); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}

func main() {
	loadEnv()
	initDB()

	todo := Todo{
		Id:          1,
		Name:        "Estudar Go",
		Description: "Vamos estudar Go",
		Done:        false,
	}

	result := SaveTodo(todo, DB)
	log.Print("Teste", result)

	fetched, err := GetTodoById(9, DB)
	if err != nil {
		log.Print("Erro ao buscar todo:", err)
	}
	log.Print("Todo encontrado:", fetched)

	allTodos, err := GetTodos(DB)
	if err != nil {
		log.Print("Erro ao buscar todos:", err)
	}
	log.Print("Todos encontrados:", allTodos)

	var teste = fetched.Toggle()
	teste2 := SaveTodo(*teste, DB)
	log.Print("Teste", teste2)

}
