package db

import (
	"path/filepath"
	"testing"

	"github.com/billadm/kernel/util"
)

func TestDbConnect(t *testing.T) {
	dbPath := filepath.Join(util.GetTestDir(), util.GetRandomString(8)+".db")
	if _, err := newDbInstance(dbPath); err != nil {
		t.Error("Error connecting to database")
	}
}

func TestOpenAndInit(t *testing.T) {
	initSqlPath := filepath.Join(util.GetConfDir(), "sql", "billadm.sql")
	dbPath := filepath.Join(util.GetTestDir(), util.GetRandomString(8)+".db")
	if err := openAndInit(dbPath, initSqlPath); err != nil {
		t.Errorf("Error execute script in database, dbPath: %s, initSqlPath: %s, err: %v", dbPath, initSqlPath, err)
	}
}
