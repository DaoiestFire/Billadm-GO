package db

import "testing"

func TestDbConnect(t *testing.T) {
	config.DatabasePath = "../../test/test.db"
	_, err := GetInstance()
	if err != nil {
		t.Error("Error connecting to database")
	}
}
