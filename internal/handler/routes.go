package handler

import "net/http"

func (h *TaskHandler) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /tasks", h.List)
	mux.HandleFunc("GET /tasks/{id}", h.GetByID)
	mux.HandleFunc("POST /tasks", h.Create)
	mux.HandleFunc("PUT /tasks/{id}", h.Update)
	mux.HandleFunc("DELETE /tasks/{id}", h.Delete)

	mux.HandleFunc("GET /tasks/trash", h.Trash)
	mux.HandleFunc("POST /tasks/{id}/restore", h.Restore)

	return mux
}
