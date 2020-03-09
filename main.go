package main

import (
	"fmt"
	"net/http"

	"github.com/ThomasHamilton2/todo-list-6/db"
	"github.com/ThomasHamilton2/todo-list-6/handler"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

func main() {
	var mySQL *db.MySQL
	var err error
	mySQL, err = db.ConnectMySQL()
	if err != nil {
		panic(err)
	} else if mySQL == nil {
		panic("mySQL is nil")
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"DELETE", "GET", "OPTIONS", "POST", "PUT"},
		// todo thomas - possibly remove for prod
		Debug: true,
	})

	mux := handler.SetUpRouting(mySQL)

	fmt.Println("http://localhost:8080")

	handler := cors.Default().Handler(mux)
	handler = c.Handler(handler)
	http.ListenAndServe(":8080", handler)
}
