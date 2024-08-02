package todoModel

import db "netix-todo-app/backend/lib/db/sqlite"

type TodoList []Todo

func (us *TodoList) FetchAll() error {
	udb := db.GormDB.Find(us)
	return udb.Error
}
