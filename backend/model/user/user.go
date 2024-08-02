package userModel

import (
	db "netix-todo-app/backend/lib/db/sqlite"
	"netix-todo-app/backend/lib/utils"
)

type User struct {
	Id       int64          `json:"id" gorm:"primaryKey"`
	Name     string         `json:"name"`
	UserName string         `json:"userName"`
	Phone    string         `json:"phone"`
	Password string         `json:"password"`
	Roles    db.StringArray `json:"roles"`
	Remark   string         `json:"remark"`
	CreateAt int64          `json:"createAt"`
	UpdateAt int64          `json:"updateAt"`
	DeleteAt int64          `json:"deleteAt"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) GetByUserName() error {
	return db.GormDB.First(u, "user_name = ? and delete_at = 0", u.UserName).Error
}
func (u *User) Get() error {
	return db.GormDB.First(u, "id = ? and delete_at = 0", u.Id).Error
}

func (u *User) Save() error {
	u.UpdateAt = utils.GetTimeStampMillSecond()
	mdb := db.GormDB.Save(u)
	return mdb.Error
}
func (u *User) Delete() error {
	u.DeleteAt = utils.GetTimeStampMillSecond()
	mdb := db.GormDB.Save(u)
	return mdb.Error
}
