package data

import "time"

type Todo struct {
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Author    string    `json:"author"`
	TimeAdded time.Time `json:"date-added"`
}

type TodoTitle struct {
	Title string `json:"title"`
}

var DeleteTodoRequest TodoTitle

var TodoRequest Todo

type ErrorResponse struct {
	Message string `json:"message"`
}

var Todos = []Todo{
	{
		"Work",
		"Wake up early and go finish my assigned tasks",
		"Mary Sharp",
		time.Now(),
	},
	{
		"Eat",
		"Eat some breakfast and take medication",
		"Mary Sharp",
		time.Now(),
	},
	{
		"Eat",
		"Eat some breakfast and take medication",
		"Mary Sharp",
		time.Now(),
	},
	{
		"Wash",
		"Wash all dirty clothes",
		"Sampson",
		time.Now(),
	},
}
