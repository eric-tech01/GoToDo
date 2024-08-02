package userSrv

import (
	"netix-todo-app/backend/lib/utils"
	userModel "netix-todo-app/backend/model/user"

	log "github.com/eric-tech01/simple-log"
)

func FindByUserName(username string) (*userModel.User, error) {
	u := &userModel.User{UserName: username}
	if err := u.GetByUserName(); err != nil {
		return nil, err
	}
	return u, nil
}

func VerifyPwd() {

}
func FindById(id int64) (*userModel.User, error) {
	u := &userModel.User{Id: id}
	if err := u.Get(); err != nil {
		return nil, err
	}
	return u, nil
}

func InitUserConfig() error {
	var users userModel.Users
	if err := users.FetchAll(); err != nil {
		log.Error(err)
		return nil
	}
	if len(users) > 0 {
		return nil
	}
	log.Info("system init admin account")
	now := utils.GetTimeStampMillSecond()
	user := userModel.User{
		Name:     "管理员",
		UserName: "admin",
		Password: utils.Md5V1("123456"),
		Roles:    []string{userModel.USER_ROLE_SUPER},
		CreateAt: now,
		UpdateAt: now,
	}
	if err := user.Save(); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
