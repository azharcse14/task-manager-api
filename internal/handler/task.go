package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/azharcse14/task-manager-api/internal/task"
)

type TaskHandler struct {
	store *task.Store
}

func NewTaskHandler(store *task.Store) *TaskHandler {
	return &TaskHandler{store: store}
}

func (h *TaskHandler) List(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.store.GetAll()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Could not fetch tasks")
		return
	}
	writeJSON(w, http.StatusOK, tasks)
}

func (h *TaskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}

	t, err := h.store.GetByID(id)
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

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newTask task.Task

	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid json body")
		return
	}

	if newTask.Title == "" {
		writeError(w, http.StatusBadRequest, "Title is required")
		return
	}

	if err := h.store.Create(&newTask); err != nil {
		writeError(w, http.StatusInternalServerError, "Could not create task")
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/tasks/%d", newTask.ID))
	writeJSON(w, http.StatusCreated, newTask)
}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}

	var updated task.Task
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	if updated.Title == "" {
		writeError(w, http.StatusBadRequest, "Title is required")
		return
	}
	updated.ID = id
	err := h.store.Update(updated)
	if errors.Is(err, sql.ErrNoRows) {
		writeError(w, http.StatusNotFound, "Task not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Could not update task")
		return
	}

	writeJSON(w, http.StatusOK, updated)
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}

	err := h.store.Delete(id)
	if errors.Is(err, sql.ErrNoRows) {
		writeError(w, http.StatusNotFound, "Task not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Could not delete task")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
