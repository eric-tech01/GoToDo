package todoModel

import (
	db "netix-todo-app/backend/lib/db/sqlite"
)

type Todo struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreateAt  int64  `json:"create_at"`
	Completed bool   `json:"completed"`
}

func (Todo) TableName() string {
	return "todo"
}

func (u *Todo) Get() error {
	mdb := db.GormDB.First(u)
	return mdb.Error
}
func (u *Todo) Save() error {
	mdb := db.GormDB.Save(u)
	return mdb.Error
}
func (u *Todo) Delete() error {
	mdb := db.GormDB.Delete(u)
	return mdb.Error
}
