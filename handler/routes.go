package handler

import (
	"net/http"

	"github.com/ThomasHamilton2/todo-list/db"
)

func SetUpRouting() *http.ServeMux {
	todoHandler := &todoHandler{
		samples: &db.Sample{},
	}

	// r := gin.Default()
	// r.Use(CORSMiddleware())
	// r.GET("/samples", todoHandler.GetSamples)

	mux := http.NewServeMux()
	mux.HandleFunc("/samples", todoHandler.GetSamples).Methods("GET")
	mux.HandleFunc("/samples", todoHandler.Delete).Methods("DELETE")
	mux.HandleFunc("/samples", todoHandler.AddTodo).Methods("POST")
	// mux.HandleFunc("/addTodo", todoHandler.addTodo)
	// mux.HandleFunc("/updateTodo", todoHandler.addTodo)

	return mux
	// return r
}

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST, PUT")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}

// 		c.Next()
// 	}
// }
