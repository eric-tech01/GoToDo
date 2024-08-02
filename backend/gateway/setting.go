package gateway

import (
	"context"
	"netix-todo-app/backend/global"
	"netix-todo-app/backend/lib/utils"
	systemModel "netix-todo-app/backend/model/system"
	systemSrv "netix-todo-app/backend/service/system"

	log "github.com/eric-tech01/simple-log"
)

// App struct
type SettingController struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewSettingController() *SettingController {
	return &SettingController{}
}

func (a *SettingController) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *SettingController) Load() global.ResponseT {
	s, err := systemSrv.Load()
	if err != nil {
		log.Error(err)
		return global.Failed(global.SYSTEM_EXCEPTION)
	}
	return global.SuccessWithData(s)
}

// 保存更新数据
func (a *SettingController) Save(t *systemModel.Setting) global.ResponseT {
	t.Id = utils.GetTimeStampMillSecond()
	if err := t.Save(); err != nil {
		log.Error(err)
		return global.Failed(global.SYSTEM_EXCEPTION)
	}
	return global.SuccessWithData(t)
}
