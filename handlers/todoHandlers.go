package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	t "github.com/ceejay1000/todo-app/data"
)

func GetAllTodos(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		errResponse, err := json.Marshal(t.ErrorResponse{Message: "Method Not Allowed"})

		if err != nil {
			w.Write([]byte("An error occured"))
		}

		w.Write(errResponse)
		return
	}

	allTodos := t.Todos
	todoJson, err := json.Marshal(allTodos)

	if err != nil {
		log.Println("Cannot encode JSON")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(todoJson)
	// fmt.Fprintf(w, string(todoJson))
	log.Println("Data sent with status code: 200")
}

func AddTodo(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		errMsg, err := json.Marshal(t.ErrorResponse{Message: "Method not supported"})

		if err != nil {
			log.Println("Cannot Marshal JSON")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		w.Write(errMsg)
		return
	}

	// json.NewDecoder(r.Body)
	todoRequest, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println("Error processing request")
	}

	todoErr := json.Unmarshal([]byte(todoRequest), &t.TodoRequest)

	if todoErr != nil {
		log.Println("Error parsing json")
		return
	}

	t.Todos = append(t.Todos, t.TodoRequest)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	todoResponse, err := json.Marshal(t.Todos)

	if err != nil {
		log.Println("Cannot parse json")
		return
	}

	w.Write(todoResponse)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {

	var updatedTodo t.Todo

	if r.Method != http.MethodPatch {
		errMsg, err := json.Marshal(t.ErrorResponse{Message: "Method not supported"})

		if err != nil {
			log.Println("Cannot Marshal JSON")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		w.Write(errMsg)
		return
	}

	newTodo, err := io.ReadAll(r.Body)

	err = json.Unmarshal(newTodo, &updatedTodo)

	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("An error occured"))
	}

	for index, todo := range t.Todos {

		if strings.EqualFold(todo.Title, updatedTodo.Title) {
			todo.Author = updatedTodo.Author
			todo.Body = updatedTodo.Body
			t.Todos[index] = todo
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Todo Updated successfully"))
			return
		}
	}

	// w.Header().Set("Content-Type", "aaplication/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("No todo with title '" + updatedTodo.Title + "' found"))

}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {

		errResponse, err := json.Marshal(t.ErrorResponse{Message: "Method not allowed"})

		if err != nil {
			log.Println("Unable to serialize data")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(errResponse)
		return
	}

	resp, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println("Unable to parse request body")
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(err.Error()))
	}

	// todoTitle := string(resp)

	err = json.Unmarshal([]byte(resp), &t.DeleteTodoRequest)

	if err == nil {
		log.Println(err)
	}

	log.Println(t.DeleteTodoRequest)

	for index, todo := range t.Todos {

		// if strings.ToLower(todoTitle) == strings.ToLower(todo.Title) {

		if strings.EqualFold(t.DeleteTodoRequest.Title, todo.Title) {
			t.Todos = append(t.Todos[:index], t.Todos[index+1:len(t.Todos)]...)
			log.Println("Deleting Todos...")

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			todosResponse, err := json.Marshal(t.Todos)

			if err != nil {
				log.Println("Unable to serialize data to JSON")
			}

			w.Write(todosResponse)
			return
		}
	}

	errResponse, err := json.Marshal(t.ErrorResponse{Message: "Todo not found"})

	if err != nil {
		log.Println("Unable to serialize data")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(errResponse)

}
