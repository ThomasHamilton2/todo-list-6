package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ThomasHamilton2/todo-list/db"
	"github.com/ThomasHamilton2/todo-list/service"
)

type todoHandler struct {
	samples *db.Sample
}

func (handler *todoHandler) GetSamples(w http.ResponseWriter, r *http.Request) {
	log.Println("---------show something---------")
	if r.Method == "Delete" {
		id, ok := r.URL.Query()["id"]
		err := service.Delete(ctx, id)
		responseOk(w, nil)
	}
	else {
		ctx := db.SetRepository(r.Context(), handler.samples)
		// setupResponse(&w, r)
		todoList, err := service.GetAll(ctx)
		if err != nil {
			responseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	
		responseOk(w, todoList)
	}
}

func (handler *todoHandler) Delete(w http.ResponseWriter, req *http.Request) {
	fmt.Println("made it inside the delete method")
	// setupResponse(&w, req)
	if (*req).Method == "OPTIONS" {
		return
	}
	id, ok := r.URL.Query()["id"]
	err := service.Delete(ctx, id)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOk(w, nil)

}

func (handler *todoHandler) AddTodo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("made it inside the add method")
	// setupResponse(&w, req)
	if (*req).Method == "OPTIONS" {
		return
	}
	var i Todo
	i.Title = "Thomas"
	i.Complete = false
	err := service.Insert(ctx, i)
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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
