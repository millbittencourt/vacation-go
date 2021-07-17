package data

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"mille.com/todo/entity"
)

func InitDatabase() {
	if _, err := os.Stat("database.db"); os.IsNotExist(err) {
		file, err := os.Create("database.db")

		if err != nil {
			log.Fatal(err.Error())
		}

		file.Close()

		log.Print("db file created")
		database, _ := sql.Open("sqlite3", "./database.db")
		log.Print("db is up")
		defer database.Close()

		createTodoTable(database)
		log.Print("table created")

		return
	}
	log.Print("db already exists")
}

func createTodoTable(db *sql.DB) {
	query := `CREATE TABLE todo (
		id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		description TEXT,
		finished INT
	);`

	statement, err := db.Prepare(query)

	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
}

func InsertTodo(todo entity.Todo) {
	db, _ := sql.Open("sqlite3", "./database.db")

	query := `INSERT INTO todo (description, finished) VALUES (?,?)`

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = statement.Exec(todo.Description, todo.Finished)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()
}

func FindTodo() []entity.Todo {
	db, _ := sql.Open("sqlite3", "./database.db")

	var todoList []entity.Todo

	row, err := db.Query("SELECT * FROM todo")
	if err != nil {
		log.Fatal(err.Error())
	}

	for row.Next() {
		var todo entity.Todo

		row.Scan(
			&todo.Id,
			&todo.Description,
			&todo.Finished,
		)
		todoList = append(todoList, todo)
	}

	defer db.Close()
	defer row.Close()

	return todoList
}

func SetTodoFinished(todoId int) {
	db, _ := sql.Open("sqlite3", "./database.db")

	query := `UPDATE todo SET finished = true WHERE id = ?`

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = statement.Exec(todoId)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
}

func DeleteTodo(todoId int) {
	db, _ := sql.Open("sqlite3", "./database.db")

	query := `DELETE FROM todo WHERE id = ?`

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = statement.Exec(todoId)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
}
