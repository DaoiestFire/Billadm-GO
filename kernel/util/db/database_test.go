package db

import "testing"

func TestDbConnect(t *testing.T) {
	Config.DatabasePath = "../../test/test.db"
	_, err := GetInstance()
	if err != nil {
		t.Error("Error connecting to database")
	}
}

func TestOpenAndInit(t *testing.T) {
	_, err := OpenAndInit("../../test/test.db", "../../test/init.sql")
	if err != nil {
		t.Error("Error execute script in database")
	}
}
