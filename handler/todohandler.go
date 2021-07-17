package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"mille.com/todo/entity"
	"mille.com/todo/service"
)

func Health(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprint(responseWriter, "Server OK 3.0")
}

func GetTodo(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		jsonBody, err := json.Marshal(service.Find())
		if err != nil {
			http.Error(responseWriter, "Error parsing todo to json", http.StatusInternalServerError)
		}

		fmt.Print(jsonBody)
		responseWriter.Write(jsonBody)
	} else {
		http.Error(responseWriter, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func InsertTodo(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(responseWriter, "Error reading request body", http.StatusInternalServerError)
		}

		var todo entity.Todo
		err = json.Unmarshal(body, &todo)
		if err != nil {
			http.Error(responseWriter, "Error parsing request body", http.StatusMethodNotAllowed)
		}

		service.Insert(todo)

	} else {
		http.Error(responseWriter, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func SetTodoFinished(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == "PUT" {
		query := request.URL.Query()
		todo_id := query.Get("todo-id")

		todoId, err := strconv.Atoi(todo_id)

		if err != nil {
			http.Error(responseWriter, "Error parsing param", http.StatusInternalServerError)
		}
		service.SetTodoFinished(todoId)

	} else {
		http.Error(responseWriter, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
