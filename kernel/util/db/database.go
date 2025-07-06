package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/billadm/util"
)

const (
	SqlFileName = "billadm.sql"
)

var db *gorm.DB

func GetInstance() *gorm.DB {
	return db
}

func Init(dbPath string) error {
	initScriptPath := filepath.Join(util.GetConfDir(), "sql", SqlFileName)
	if err := openAndInit(dbPath, initScriptPath); err != nil {
		return err
	}

	ins, err := newDbInstance(dbPath)
	if err != nil {
		return err
	}
	db = ins
	return nil
}

func newDbInstance(dbPath string) (*gorm.DB, error) {
	var err error
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logrus.Errorf("连接数据库失败, db path: %s, err: %v", dbPath, err)
		return nil, fmt.Errorf("连接数据库失败, db path: %s, err: %v", dbPath, err)
	}
	logrus.Infof("连接数据库成功, db path: %s", dbPath)
	return db, nil
}

// 打开数据库并执行初始化脚本
func openAndInit(dbPath, initScriptPath string) error {
	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("打开数据库失败: %w", err)
	}

	if err = database.Ping(); err != nil {
		return fmt.Errorf("数据库连接验证失败: %w", err)
	}

	script, err := os.ReadFile(initScriptPath)
	if err != nil {
		return fmt.Errorf("读取初始化脚本失败: %w", err)
	}

	if err = executeInitScript(database, string(script)); err != nil {
		database.Close()
		return fmt.Errorf("执行初始化脚本失败: %w", err)
	}

	return nil
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
