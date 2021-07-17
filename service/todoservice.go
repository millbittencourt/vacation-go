package service

import (
	"mille.com/todo/entity"
)

var todoList []entity.Todo
var id int

func Insert(todo entity.Todo) {
	id++
	todo.Id = id
	todoList = append(todoList, todo)
}

func Find() []entity.Todo {
	return todoList
}

func SetTodoFinished(todoId int) {
	for i, todo := range todoList {
		if todoId == todo.Id {
			todoList[i].Finished = true
		}
	}
}

func DeleteTodo(todoId int) {
	var newTodoList []entity.Todo
	for _, todo := range todoList {
		if todoId != todo.Id {
			newTodoList = append(newTodoList, todo)
		}
	}
	if newTodoList != nil {
		// todoList = nil
		todoList = newTodoList
	}
}
