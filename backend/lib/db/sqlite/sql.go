package sqlite

import (
	log "github.com/eric-tech01/simple-log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func InitDB(name string, args ...interface{}) error {
	var err error
	dbPath := "./" + name
	GormDB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Error(err)
		return err
	}
	for _, a := range args {
		err := GormDB.AutoMigrate(&a)
		if err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}
