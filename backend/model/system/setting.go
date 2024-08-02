package systemModel

import (
	db "netix-todo-app/backend/lib/db/sqlite"
)

type Setting struct {
	Id         int64 `json:"id" gorm:"primaryKey"`
	AuthOpened bool  `json:"authOpened"` //是否开启验证
	IsShow     bool  `json:"isShow"`     //隐藏显示
}

func (Setting) TableName() string {
	return "setting"
}

func (u *Setting) Get() error {
	u.Id = 1
	return db.GormDB.First(u, "id=?", u.Id).Error
}
func (u *Setting) Save() error {
	u.Id = 1
	return db.GormDB.Save(u).Error
}
