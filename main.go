package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
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
var mu sync.Mutex

func main() {
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

	mu.Lock()
	defer mu.Unlock()

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	writeError(w, http.StatusNotFound, "Task not found")
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

	mu.Lock()
	defer mu.Unlock()

	for i, task := range tasks {
		if task.ID == id {
			updated.ID = id
			tasks[i] = updated
			writeJSON(w, http.StatusOK, updated)
			return
		}
	}
	writeError(w, http.StatusNotFound, "Task not found")
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

	mu.Lock()
	defer mu.Unlock()

	newTask.ID = nextID
	nextID++

	tasks = append(tasks, newTask)

	writeJSON(w, http.StatusCreated, newTask)
}

func tasksByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for _, task := range tasks {
		if task.ID == id {
			writeJSON(w, http.StatusOK, task)
			return
		}
	}
	writeError(w, http.StatusBadRequest, "Invalid task ID")
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

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
