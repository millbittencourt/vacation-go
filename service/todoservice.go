package service

import (
	"mille.com/todo/data"
	"mille.com/todo/entity"
)

var todoList []entity.Todo
var id int

func Insert(todo entity.Todo) {
	data.InsertTodo(todo)
}

func Find() []entity.Todo {
	return data.FindTodo()
}

func SetTodoFinished(todoId int) {
	data.SetTodoFinished(todoId)
}

func DeleteTodo(todoId int) {
	data.DeleteTodo(todoId)
}
