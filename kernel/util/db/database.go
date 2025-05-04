package db

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
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

// OpenAndInit 打开数据库并执行初始化脚本
func OpenAndInit(dbPath string, initScriptPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("打开数据库失败: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("数据库连接验证失败: %w", err)
	}

	script, err := os.ReadFile(initScriptPath)
	if err != nil {
		return nil, fmt.Errorf("读取初始化脚本失败: %w", err)
	}

	if err = executeInitScript(db, string(script)); err != nil {
		db.Close()
		return nil, fmt.Errorf("执行初始化脚本失败: %w", err)
	}

	return db, nil
}

// executeInitScript 执行初始化 SQL 脚本
func executeInitScript(db *sql.DB, script string) error {
	statements := strings.Split(script, ";")

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	for _, stmt := range statements {
		cleaned := strings.TrimSpace(stmt)
		if cleaned == "" {
			continue
		}
		if _, err = tx.Exec(cleaned); err != nil {
			return fmt.Errorf("执行 SQL 失败: %s\n错误: %w", cleaned, err)
		}
	}

	// 提交事务
	return tx.Commit()
}
