package main

import (
	"fmt"
	"log"
	"net/http"

	t "github.com/ceejay1000/todo-app/handlers"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the todo api")
}

func main() {

	h := new(handler)
	router := http.NewServeMux()

	router.Handle("/", h)

	router.HandleFunc("/get-todos", t.GetAllTodos)

	router.HandleFunc("/add-todo", t.AddTodo)

	router.HandleFunc("/delete-todo", t.DeleteTodo)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	log.Println("Server running on port 9000")
	if err := server.ListenAndServe(); err != nil {
		log.Println(err.Error())
	}
}
