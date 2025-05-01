package sqlite

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func OpenDB() {
	var err error
	DB, err = sql.Open("sqlite", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

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
