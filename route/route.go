package route

import (
	"net/http"

	"mille.com/todo/handler"
)

func RegisterRoute() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Health)
	mux.HandleFunc("/todo", handler.GetTodo)
	mux.HandleFunc("/todo/add", handler.InsertTodo)
	mux.HandleFunc("/todo/set-finished", handler.SetTodoFinished)

	return mux
}
