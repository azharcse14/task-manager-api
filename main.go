package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "modernc.org/sqlite"

	"github.com/azharcse14/task-manager-api/internal/handler"
	"github.com/azharcse14/task-manager-api/internal/task"
)

func main() {
	db, err := sql.Open("sqlite", "tasks.db")
	if err != nil {
		log.Fatal("Cannot open database:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	if err := task.CreateSchema(db); err != nil {
		log.Fatal("Cannot create schema:", err)
	}

	store := task.NewStore(db)
	taskHandler := handler.NewTaskHandler(store)

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", taskHandler.Routes()); err != nil {
		log.Fatal("Server failed:", err)
	}
}
