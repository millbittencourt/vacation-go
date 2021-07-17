package service

import (
	"mille.com/todo/data"
	"mille.com/todo/entity"
)

var todoList []entity.Todo
var id int

func Insert(todo entity.Todo) {
	data.InsertTodo(todo)
	// id++
	// todo.Id = id
	// todoList = append(todoList, todo)
}

func Find() []entity.Todo {
	return data.FindTodo()
	// return todoList
}

func SetTodoFinished(todoId int) {
	data.SetTodoFinished(todoId)
	// for _, todo := range todoList {
	// if todoId == todo.Id {
	// todoList[i].Finished = true
	// }
	// }
}
