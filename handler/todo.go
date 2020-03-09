package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ThomasHamilton2/todo-list-6/db"
	"github.com/ThomasHamilton2/todo-list-6/schema"
	"github.com/ThomasHamilton2/todo-list-6/service"
)

type todoHandler struct {
	samples *db.MySQL
}

func (handler *todoHandler) Todo(w http.ResponseWriter, r *http.Request) {
	switch (*r).Method {
	case "GET":
		TodoGet(w, r, handler)
	case "POST":
		TodoPost(w, r, handler)
	case "PUT":
		TodoPut(w, r, handler)
	case "DELETE":
		TodoDelete(w, r, handler)
	default:
		responseError(w, http.StatusMethodNotAllowed, "Status method not allowed")
	}
}

//TodoGet handles the Todo GET request
func TodoGet(w http.ResponseWriter, r *http.Request, handler *todoHandler) {
	ctx := db.SetRepository(r.Context(), handler.samples)
	// setupResponse(&w, r)
	todoList, err := service.GetAll(ctx)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOk(w, todoList)
}

//TodoDelete handles the Todo DELETE request
func TodoDelete(w http.ResponseWriter, r *http.Request, handler *todoHandler) {
	ctx := db.SetRepository(r.Context(), handler.samples)

	idStr := r.URL.Query().Get("id")
	id, strErr := strconv.Atoi(idStr)
	if strErr != nil {
		responseError(w, http.StatusInternalServerError, strErr.Error())
		return
	}

	err := service.Delete(ctx, id)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseOk(w, nil)
}

//TodoPost handles the Todo POST request
func TodoPost(w http.ResponseWriter, r *http.Request, handler *todoHandler) {
	ctx := db.SetRepository(r.Context(), handler.samples)

	var p schema.Todo
	decodeErr := json.NewDecoder(r.Body).Decode(&p)
	if decodeErr != nil {
		http.Error(w, decodeErr.Error(), http.StatusBadRequest)
		return
	}

	id, err := service.Insert(ctx, &p)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOk(w, id)
}

//TodoPut handles the Todo PUT request
func TodoPut(w http.ResponseWriter, r *http.Request, handler *todoHandler) {
	ctx := db.SetRepository(r.Context(), handler.samples)

	var p schema.Todo
	decodeErr := json.NewDecoder(r.Body).Decode(&p)
	if decodeErr != nil {
		http.Error(w, decodeErr.Error(), http.StatusBadRequest)
		return
	}

	err := service.Update(ctx, &p)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOk(w, nil)
}

func responseOk(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(body)
}

func responseError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	body := map[string]string{
		"error": message,
	}
	json.NewEncoder(w).Encode(body)
}
