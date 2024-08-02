package systemSrv

import (
	systemModel "netix-todo-app/backend/model/system"

	log "github.com/eric-tech01/simple-log"
)

// 加载setting
func Load() (systemModel.Setting, error) {
	setting := systemModel.Setting{}
	if err := setting.Get(); err != nil {
		log.Error(err)
		return setting, err
	}
	return setting, nil
}

func UpdateSetting(s *systemModel.Setting) error {
	if err := s.Save(); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
