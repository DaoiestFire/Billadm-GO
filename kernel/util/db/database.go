package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/billadm/kernel/logger"
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
	Config DatabaseConfig
)

func GetInstance() (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}

	var err error
	once.Do(func() {
		db, err = gorm.Open(sqlite.Open(Config.DatabasePath), &gorm.Config{})
		if err != nil {
			logger.Error("连接数据库失败, db path: %s, err: %v", Config.DatabasePath, err)
		}
	})

	if db != nil {
		logger.Info("连接数据库成功, db path: %s", Config.DatabasePath)
		return db, nil
	}
	return nil, err
}

// OpenAndInit 打开数据库并执行初始化脚本
func OpenAndInit(dbPath string, initScriptPath string) (*sql.DB, error) {
	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("打开数据库失败: %w", err)
	}

	if err = database.Ping(); err != nil {
		return nil, fmt.Errorf("数据库连接验证失败: %w", err)
	}

	script, err := os.ReadFile(initScriptPath)
	if err != nil {
		return nil, fmt.Errorf("读取初始化脚本失败: %w", err)
	}

	if err = executeInitScript(database, string(script)); err != nil {
		database.Close()
		return nil, fmt.Errorf("执行初始化脚本失败: %w", err)
	}

	return database, nil
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
