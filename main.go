package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "modernc.org/sqlite"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite", "tasks.db")
	if err != nil {
		log.Fatal("Cannot open database:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	createTable()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", homeHandler)
	mux.HandleFunc("GET /health", healthHandler)

	mux.HandleFunc("GET /tasks", tasksHandler)
	mux.HandleFunc("GET /tasks/{id}", tasksByIDHandler)
	mux.HandleFunc("POST /tasks", createTaskHandler)
	mux.HandleFunc("PUT /tasks/{id}", updateTaskHandler)
	mux.HandleFunc("DELETE /tasks/{id}", deleteTaskHandler)

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", mux)
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}

	result, err := db.Exec("DELETE FROM tasks WHERE id=?", id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Could not delete task")
		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Could not verify delete")
		return
	}

	if rowsAffected == 0 {
		writeError(w, http.StatusNotFound, "Task not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}

	var updated Task
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	if updated.Title == "" {
		writeError(w, http.StatusBadRequest, "Title is required")
		return
	}

	result, err := db.Exec(
		"UPDATE tasks SET title = ?, done = ? WHERE id = ?", updated.Title, updated.Done, id)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Could not update task")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Could not verify update")
		return
	}

	if rowsAffected == 0 {
		writeError(w, http.StatusNotFound, "Task not found")
		return
	}

	updated.ID = id

	writeJSON(w, http.StatusOK, updated)
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	var newTask Task

	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid json body")
		return
	}

	if newTask.Title == "" {
		writeError(w, http.StatusBadRequest, "Title is required")
		return
	}

	result, err := db.Exec(
		"INSERT INTO tasks (title, done) VALUES (?, ?)", newTask.Title, newTask.Done)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Could not create task")
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Could not read new task ID")
		return
	}

	newTask.ID = int(id)

	w.Header().Set("Location", fmt.Sprintf("/tasks/%d", newTask.ID))
	writeJSON(w, http.StatusCreated, newTask)
}

func tasksByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}

	var t Task
	err := db.QueryRow(
		"SELECT id, title, done FROM tasks WHERE id = ?", id).Scan(&t.ID, &t.Title, &t.Done)

	if errors.Is(err, sql.ErrNoRows) {
		writeError(w, http.StatusNotFound, "Task not found")
		return
	}

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Could not fetch task")
		return
	}
	writeJSON(w, http.StatusOK, t)
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, done FROM tasks ORDER BY id ASC")
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Could not fetch tasks")
		return
	}
	defer rows.Close()

	tasks := []Task{}

	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Done); err != nil {
			writeError(w, http.StatusInternalServerError, "Could not read tasks")
			return
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		writeError(w, http.StatusInternalServerError, "Could not read tasks")
		return
	}

	writeJSON(w, http.StatusOK, tasks)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Task Manager API")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is Alive")
}

func parseID(w http.ResponseWriter, r *http.Request) (int, bool) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid task ID")
		return 0, false
	}
	return id, true
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func createTable() {
	query := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		done BOOLEAN NOT NULL DEFAULT 0
	)`

	if _, err := db.Exec(query); err != nil {
		log.Fatal("Cannot create table:", err)
	}
}
