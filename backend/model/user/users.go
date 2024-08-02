package userModel

import db "netix-todo-app/backend/lib/db/sqlite"

type Users []User

func (us *Users) FetchAll() error {
	udb := db.GormDB.Find(us, "delete_at=0")
	return udb.Error
}
