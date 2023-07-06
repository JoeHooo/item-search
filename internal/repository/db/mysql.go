package db

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"item-search/internal/repository"
	"item-search/pkg/config"
	"time"
)

func Init() {
	sqlConf := config.Conf.MySQL
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		sqlConf.Username,
		sqlConf.Password,
		sqlConf.Host,
		sqlConf.Port,
		sqlConf.Database,
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("gorm init error, %s", err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		log.Printf("gorm init error: %s", err)
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)
	repository.Db = db
}
