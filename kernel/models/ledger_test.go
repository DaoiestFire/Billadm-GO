package models

import (
	"path/filepath"
	"testing"

	"github.com/billadm/util"
	"github.com/billadm/util/db"
)

func TestLedger_CRUD(t *testing.T) {
	dbPath := filepath.Join(util.GetTestDir(), util.GetRandomString(8)+".db")
	err := db.Init(dbPath)
	if err != nil {
		t.Errorf("Error execute script in database, dbPath: %s, err: %v", dbPath, err)
		return
	}

	DB := db.GetInstance()

	// 插入
	insertData := &Ledger{
		ID:   util.GetUUID(),
		Name: "test_ledger",
	}
	ret := DB.Create(insertData)
	t.Logf("insert ret: %+v", ret)

	// 查询
	var queryData Ledger
	DB.Where("name = ?", "test_ledger").First(&queryData)
	t.Logf("inserted data: %+v", queryData)
	if queryData.Name != "test_ledger" {
		t.Errorf("Name not correct, got: %s, want: %s.", queryData.Name, "test_ledger")
		return
	}

	// 修改
	id := queryData.ID
	var modifyData Ledger
	DB.Model(&Ledger{}).Where("id = ?", id).Update("name", "new_ledger")
	DB.Where("id = ?", id).First(&modifyData)
	t.Logf("modified data: %+v", modifyData)
	if modifyData.Name != "new_ledger" {
		t.Errorf("Name not correct, got: %s, want: %s.", modifyData.Name, "new_ledger")
		return
	}

	// 删除
	var deleteData Ledger
	DB.Where("id = ?", id).Delete(&Ledger{})
	DB.Where("id = ?", id).First(&deleteData)
	t.Logf("deleted data: %+v", deleteData)
	if deleteData.Name == "new_ledger" {
		t.Errorf("Name not correct, got: %s, want: %s.", deleteData.Name, "")
		return
	}
}
