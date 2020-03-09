package handler

import (
	"net/http"

	"github.com/ThomasHamilton2/todo-list-6/db"
)

//SetUpRouting sets up routing
func SetUpRouting(mySQL *db.MySQL) *http.ServeMux {
	todoHandler := &todoHandler{
		repository: mySQL,
	}

	mux := http.NewServeMux()
	//only 1 endpoint, handles GET, POST, PUT, and DELETE
	mux.HandleFunc("/todo", todoHandler.Todo)

	return mux
}
