package db

import (
	"fmt"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	DatabasePath string
}

var (
	once   sync.Once
	db     *gorm.DB
	err    error
	config DatabaseConfig
)

func GetInstance() (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}
	once.Do(func() {
		var err error
		db, err = gorm.Open(sqlite.Open(config.DatabasePath), &gorm.Config{})
		if err != nil {
			fmt.Println("连接数据库失败：", err)
		}
	})
	if db != nil {
		fmt.Println("连接数据库成功")
		return db, nil
	}
	return nil, err
}
