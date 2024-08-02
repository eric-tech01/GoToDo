package backend

import (
	"netix-todo-app/backend/components/resource"
	"netix-todo-app/backend/lib/db/sqlite"
	systemModel "netix-todo-app/backend/model/system"
	todoModel "netix-todo-app/backend/model/todo"
	userModel "netix-todo-app/backend/model/user"

	log "github.com/eric-tech01/simple-log"
)

type Engine struct {
}

func (e *Engine) Startup(fns ...func() error) error {
	for _, fn := range fns {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}
func (e *Engine) InitLog() error {
	log.SetOptions("./logs/slice.app.log", &log.Option{
		MaxSizeInMB: 10,
		MaxBackups:  10,
		Compress:    true,
		LocalTime:   true,
	})
	return nil
}

func (e *Engine) LoadResource() error {
	return resource.Load()
}

func (e *Engine) InitDB() error {
	return sqlite.InitDB("todo.db", userModel.User{}, todoModel.Todo{}, systemModel.Setting{})
}

func (e *Engine) InitSystemConfig() error {
	log.Info("start init system init config")
	//TODO：初始化系统账号
	//TODO：初始化系统配置：例如禁用登录，采用默认super账号登录

	//初始化系统配置
	s := systemModel.Setting{Id: 1}
	if err := s.Get(); err == nil {
		//已经初始化了
		log.Info("repeat init setting, ignore")
		return nil
	} else {
		s.AuthOpened = true
		s.IsShow = true
		if err := s.Save(); err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}
