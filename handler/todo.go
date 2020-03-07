package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ThomasHamilton2/todo-list/db"
	"github.com/ThomasHamilton2/todo-list/service"
)

type todoHandler struct {
	samples *db.Sample
}

func (handler *todoHandler) GetSamples(w http.ResponseWriter, r *http.Request) {
	ctx := db.SetRepository(r.Context(), handler.samples)
	enableCors(&w)
	todoList, err := service.GetAll(ctx)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOk(w, todoList)
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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
