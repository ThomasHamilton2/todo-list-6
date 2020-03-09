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
	//first, connect to MySQL db
	var mySQL *db.MySQL
	var err error
	mySQL, err = db.ConnectMySQL()
	if err != nil {
		panic(err)
	} else if mySQL == nil {
		panic("mySQL is nil")
	}

	//set cors to allow requests to come from angular
	//todo thomas - find better option than accepting all origins
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"DELETE", "GET", "OPTIONS", "POST", "PUT"},
		// todo thomas - possibly remove for prod
		Debug: true,
	})

	//set up routes and attach(?) db instance
	mux := handler.SetUpRouting(mySQL)

	fmt.Println("http://localhost:8080")

	//attach cors options from above to the handler
	handler := cors.Default().Handler(mux)
	handler = c.Handler(handler)

	//serve on :8080
	http.ListenAndServe(":8080", handler)
}
