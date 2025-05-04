package db

import (
	"path/filepath"
	"testing"

	"github.com/billadm/kernel/util"
)

func TestDbConnect(t *testing.T) {
	Config.DatabasePath = filepath.Join(util.GetTestDir(), util.GetRandomString(8)+".db")
	if _, err := GetInstance(); err != nil {
		t.Error("Error connecting to database")
	}
}

func TestOpenAndInit(t *testing.T) {
	initSqlPath := filepath.Join(util.GetTestDir(), "init.sql")
	dbPath := filepath.Join(util.GetTestDir(), util.GetRandomString(8)+".db")
	if _, err := OpenAndInit(dbPath, initSqlPath); err != nil {
		t.Errorf("Error execute script in database, dbPath: %s, initSqlPath: %s, err: %v", dbPath, initSqlPath, err)
	}
}
