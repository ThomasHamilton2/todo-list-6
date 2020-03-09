package handler

import (
	"net/http"

	"github.com/ThomasHamilton2/todo-list-6/db"
)

//SetUpRouting sets up routing
func SetUpRouting(mySQL *db.MySQL) *http.ServeMux {
	todoHandler := &todoHandler{
		samples: mySQL,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/todo", todoHandler.Todo)

	return mux
}
