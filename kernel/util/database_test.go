package util

import (
	"path/filepath"
	"testing"
)

func TestDbConnect(t *testing.T) {
	dbPath := filepath.Join(GetRootDir(), GetRandomString(8)+".db")
	if _, err := newDbInstance(dbPath); err != nil {
		t.Error("Error connecting to database")
	}
}

func TestOpenAndInit(t *testing.T) {
	dbPath := filepath.Join(GetRootDir(), GetRandomString(8)+".db")
	if err := openAndInit(dbPath); err != nil {
		t.Errorf("Error execute script in database, dbPath: %s, err: %v", dbPath, err)
	}
}
