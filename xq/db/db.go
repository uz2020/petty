package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

func InitDb(user, passwd, addr, db string) (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=3s",
		user, passwd, addr, db)

	defaultLogger := logger.Default.LogMode(logger.Silent)
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   defaultLogger,
		SkipDefaultTransaction:                   true,
	}

	if db, err := gorm.Open(mysql.Open(dsn), gormConfig); err == nil {
		err = db.Use(dbresolver.Register(dbresolver.Config{}).
			SetConnMaxLifetime(1 * time.Hour).
			SetMaxIdleConns(100).
			SetMaxOpenConns(1000))
		if err != nil {
			panic(err)
		}
		return db, nil
	} else {
		return nil, err
	}
}
