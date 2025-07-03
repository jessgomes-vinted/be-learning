package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

var watchCourse = "Watch Go course"
var buildApp = "Build a Go app"
var reward = "Reward myself with a snack"

var taskItems = []string{watchCourse, buildApp, reward}

var greeting = "Hello, User!"

func main() {
	fmt.Println("### My To-Do List ###")

	// Create a new database
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Error closing database: ", err)
		} else {
			fmt.Println("Database connection closed successfully.")
		}
	}(db)

	// Create a table for tasks
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// Insert tasks into the table
	for _, task := range taskItems {
		_, err = db.Exec("INSERT INTO tasks (name) VALUES (?)", task)
		if err != nil {
			log.Fatal(err)
		}
	}

	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}

}

func showTasks(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, greeting)

	for _, task := range taskItems {
		_, _ = fmt.Fprintln(w, task)
	}

}

func helloUser(w http.ResponseWriter, _ *http.Request) {

	_, _ = fmt.Fprintf(w, greeting)

}
