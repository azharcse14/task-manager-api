package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks = []Task{
	{ID: 1, Title: "Learn Go basics", Done: true},
	{ID: 2, Title: "Build an API", Done: false},
}

var nextID = 3

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", homeHandler)
	mux.HandleFunc("GET /health", healthHandler)

	mux.HandleFunc("GET /tasks", tasksHandler)
	mux.HandleFunc("GET /tasks/{id}", tasksByIDHandler)
	mux.HandleFunc("POST /tasks", createTaskHandler)

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", mux)
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

	newTask.ID = nextID
	nextID++

	tasks = append(tasks, newTask)

	writeJSON(w, http.StatusCreated, newTask)
}

func tasksByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	for _, task := range tasks {
		if task.ID == id {
			writeJSON(w, http.StatusOK, task)
			return
		}
	}
	writeError(w, http.StatusBadRequest, "Invalid task ID")
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, tasks)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Task Manager API")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is Alive")
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
