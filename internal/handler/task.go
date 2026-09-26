package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/azharcse14/task-manager-api/internal/task"
)

type TaskHandler struct {
	store *task.Store
}

func NewTaskHandler(store *task.Store) *TaskHandler {
	return &TaskHandler{store: store}
}

// pageInfo: পাতার হিসাব, যেটা উত্তরের সাথে যাবে
type pageInfo struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

// listResponse: GET /tasks এর পুরো উত্তর দেখতে কেমন হবে
type listResponse struct {
	Data []task.Task `json:"data"`
	Meta pageInfo    `json:"meta"`
}

func (h *TaskHandler) List(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	// ১. URL থেকে query পড়ো
	page, err := intQuery(q, "page", 1)
	if err != nil {
		writeError(w, http.StatusBadRequest, "page must be a number")
		return
	}
	perPage, err := intQuery(q, "per_page", 10)
	if err != nil {
		writeError(w, http.StatusBadRequest, "per_page must be a number")
		return
	}
	done, err := boolQuery(q, "done")
	if err != nil {
		writeError(w, http.StatusBadRequest, "done must be true or false")
		return
	}

	// ২. মানগুলো ঠিক সীমার মধ্যে রাখো
	page = max(page, 1)
	perPage = min(max(perPage, 1), 100)
	offset := (page - 1) * perPage

	// ৩. একটা header পড়ো (শুধু দেখার জন্য লগে লিখছি)
	slog.Info("list tasks",
		"page", page,
		"per_page", perPage,
		"user_agent", r.Header.Get("User-Agent"),
	)

	// ৪. ডাটাবেস থেকে আনো
	tasks, total, err := h.store.List(r.Context(), task.ListFilter{
		Done:   done,
		Limit:  perPage,
		Offset: offset,
	})
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Could not fetch tasks")
		return
	}

	// ৫. উত্তরে header লেখো, তারপর JSON পাঠাও
	w.Header().Set("X-Total-Count", strconv.Itoa(total))
	writeJSON(w, http.StatusOK, listResponse{
		Data: tasks,
		Meta: pageInfo{
			Page:       page,
			PerPage:    perPage,
			Total:      total,
			TotalPages: (total + perPage - 1) / perPage,
		},
	})
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
