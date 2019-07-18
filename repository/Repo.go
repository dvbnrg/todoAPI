package repository

import (
	"fmt"

	model "github.com/dvbnrg/todoAPI/model"
)

var currentId int

var Todos model.Todos

func RepoFindTodo(id int) model.Todo {
	var todo model.Todo
	
	db := dbConn()
	
	results, err := db.Query("SELECT * FROM todos WHERE id=?", id)
	if err != nil {
		panic(err)
	}

	for results.Next(){
		results.Scan(&todo.Id, &todo.Name, &todo.Completed)
	}
	defer db.Close()
	return todo
}

func RepoCreateTodo(t model.Todo) model.Todo {
	currentId += 1
	t.Id = currentId
	Todos = append(Todos, t)
	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range Todos {
		if t.Id == id {
			Todos = append(Todos[:i], Todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
