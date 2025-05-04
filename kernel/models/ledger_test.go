package models

import (
	"path/filepath"
	"testing"

	"github.com/billadm/kernel/util"
	"github.com/billadm/kernel/util/db"
)

func TestLedger_CRUD(t *testing.T) {
	initSqlPath := filepath.Join(util.GetTestDir(), "init.sql")
	dbPath := filepath.Join(util.GetTestDir(), util.GetRandomString(8)+".db")
	_, err := db.OpenAndInit(dbPath, initSqlPath)
	if err != nil {
		t.Errorf("Error execute script in database, sqlPath: %s, dbPath: %s, err: %v", initSqlPath, dbPath, err)
		return
	}
	db.Config.DatabasePath = dbPath
	DB, err := db.GetInstance()
	if err != nil {
		t.Errorf("Error connecting to database, dbPath: %s, error: %v", dbPath, err)
		return
	}

	// 插入
	insertData := &Ledger{
		ID:     util.GetUUID(),
		UserID: util.GetUUID(),
		Name:   "test_ledger",
	}
	DB.Create(insertData)

	// 查询
	var queryData Ledger
	DB.First(&queryData, "name = ?", "test_ledger")
	t.Logf("inserted data: %+v", queryData)
	if queryData.Name != "test_ledger" {
		t.Errorf("Name not correct, got: %s, want: %s.", queryData.Name, "test_ledger")
		return
	}

	// 修改
	id := queryData.ID
	var modifyData Ledger
	DB.Model(&Ledger{}).Where("id = ?", id).Update("name", "new_ledger")
	DB.First(&modifyData, "id = ?", id)
	t.Logf("modified data: %+v", modifyData)
	if modifyData.Name != "new_ledger" {
		t.Errorf("Name not correct, got: %s, want: %s.", modifyData.Name, "new_ledger")
		return
	}

	// 删除
	var deleteData Ledger
	DB.Delete(&deleteData, "id = ?", id)
	DB.First(&deleteData, "id = ?", id)
	t.Logf("deleted data: %+v", deleteData)
	if deleteData.Name == "new_ledger" {
		t.Errorf("Name not correct, got: %s, want: %s.", deleteData.Name, "")
		return
	}
}
