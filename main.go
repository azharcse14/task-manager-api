package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", homeHandler)
	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("GET /tasks", tasksHandler)

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", mux)
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := []Task{
		{ID: 1, Title: "Learn Go basics", Done: true},
		{ID: 2, Title: "Build an API", Done: false},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Task Manager API")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is Alive")
}
