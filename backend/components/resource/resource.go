package resource

import (
	"netix-todo-app/backend/global"
	"time"

	log "github.com/eric-tech01/simple-log"
)

/**
* TODO: 安全需要的话需要在此进行资源加载鉴权
* 系统启动之后，加载本地的资源文件（资源文件为硬件信息，提前私钥加密好放置在这里的）
* TODO: 暂时只支持唯一key的校验
**/
func Load() error {

	log.Info("resource load")

	global.HARD_RESOURCE.BoardModel = "metix"
	global.HARD_RESOURCE.License = global.License{
		Key:      "deviceKey",
		Expired:  time.Now().Add(time.Duration(24) * time.Hour).UnixMilli(),
		CreateAt: time.Now().UnixMilli(),
	}
	return nil

}
