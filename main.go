package main

import (
	"fmt"
	"net/http"

	"github.com/ThomasHamilton2/todo-list/handler"
	"github.com/ThomasHamilton2/todo-list/mysqlpk"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

func main() {

	// db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/Todo_db")
	// if err != nil {
	// 	panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	// }
	// fmt.Println("made it this far")

	// id := 1
	// var col string
	// sqlStatement := `SELECT Title FROM Todo WHERE ID=?`
	// row := db.QueryRow(sqlStatement, id)
	// err2 := row.Scan(&col)
	// if err2 != nil {
	// 	if err2 == sql.ErrNoRows {
	// 		fmt.Println("Zero rows found")
	// 	} else {
	// 		panic(err2)
	// 	}
	// }
	// fmt.Println("col" + col)

	var mySQL *mysqlpk.MySQL
	var err error
	mySQL, err = mysqlpk.ConnectMySQL()
	if err != nil {
		panic(err)
	} else if mySQL == nil {
		panic("mySQL is nil")
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"DELETE", "GET", "OPTIONS", "POST", "PUT"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	mux := handler.SetUpRouting(mySQL)

	fmt.Println("http://localhost:8080")

	handler := cors.Default().Handler(mux)
	handler = c.Handler(handler)
	http.ListenAndServe(":8080", handler)
}
