package todoSrv

import (
	todoModel "netix-todo-app/backend/model/todo"
	"sort"

	log "github.com/eric-tech01/simple-log"
)

func FetchAll() todoModel.TodoList {
	var todos todoModel.TodoList
	if err := todos.FetchAll(); err != nil {
		log.Error(err)
	}
	//sort一下
	sort.Slice(todos, func(i, j int) bool {
		if todos[i].Completed {
			if todos[j].Completed {
				return todos[i].Id > todos[j].Id
			}
			return false
		}
		return todos[i].Id > todos[j].Id
	})
	return todos
}
