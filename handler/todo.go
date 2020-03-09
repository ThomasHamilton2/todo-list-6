package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ThomasHamilton2/todo-list-6/db"
	"github.com/ThomasHamilton2/todo-list-6/schema"
)

type todoHandler struct {
	repository *db.MySQL
}

//Todo - all requests from /Todo are sent here, then sent to different
//functions depending on Request Method
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
	todoList, err := handler.repository.GetAll()
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOk(w, todoList)
}

//TodoDelete handles the Todo DELETE request
func TodoDelete(w http.ResponseWriter, r *http.Request, handler *todoHandler) {
	//ID of Todo to remove is stored in the URL
	//Retrieve and convert to int
	idStr := r.URL.Query().Get("id")
	id, strErr := strconv.Atoi(idStr)
	if strErr != nil {
		responseError(w, http.StatusInternalServerError, strErr.Error())
		return
	}

	err := handler.repository.Delete(id)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseOk(w, nil)
}

//TodoPost handles the Todo POST request
func TodoPost(w http.ResponseWriter, r *http.Request, handler *todoHandler) {
	//Todo data is stored in the request body
	//Decode into Todo object and insert into DB
	var p schema.Todo
	decodeErr := json.NewDecoder(r.Body).Decode(&p)
	if decodeErr != nil {
		http.Error(w, decodeErr.Error(), http.StatusBadRequest)
		return
	}

	id, err := handler.repository.Insert(&p)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOk(w, id)
}

//TodoPut handles the Todo PUT request
func TodoPut(w http.ResponseWriter, r *http.Request, handler *todoHandler) {
	//Todo data is stored in the request body
	//Decode into Todo object and update in DB
	var p schema.Todo
	decodeErr := json.NewDecoder(r.Body).Decode(&p)
	if decodeErr != nil {
		http.Error(w, decodeErr.Error(), http.StatusBadRequest)
		return
	}

	err := handler.repository.Update(&p)
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
