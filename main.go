package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ThomasHamilton2/todo-list/handler"
)

func main() {
	mux := handler.SetUpRouting()

	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
