package models

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/billadm/util"
	"github.com/billadm/util/db"
)

func TestTR_CRUD(t *testing.T) {
	dbPath := filepath.Join(util.GetRootDir(), util.GetRandomString(8)+".db")
	err := db.Init(dbPath)
	if err != nil {
		t.Errorf("Error execute script in database, dbPath: %s, err: %v", dbPath, err)
		return
	}

	DB := db.GetInstance()

	// 插入
	id := util.GetUUID()
	insertData := &TransactionRecord{
		TransactionID:   id,
		LedgerID:        util.GetUUID(),
		Price:           114.514,
		TransactionType: "",
		Category:        "",
		Description:     "this is a test tr",
		TransactionAt:   time.Now().Unix(),
	}
	DB.Create(insertData)

	// 查询
	var queryData TransactionRecord
	DB.Where("transaction_id = ?", id).First(&queryData)
	t.Logf("inserted data: %+v", queryData)
	if queryData.TransactionID != id {
		t.Errorf("transaction id should be %s, but %s", id, queryData.TransactionID)
		return
	}

	// 修改
	var modifyData TransactionRecord
	DB.Model(&TransactionRecord{}).Where("transaction_id = ?", id).Update("price", 3442.5)
	DB.Where("transaction_id = ?", id).First(&modifyData)
	t.Logf("modified data: %+v", modifyData)

	// 删除
	var deleteData TransactionRecord
	DB.Where("transaction_id = ?", id).Delete(&TransactionRecord{})
	DB.Where("transaction_id = ?", id).First(&deleteData)
	t.Logf("deleted data: %+v", deleteData)
	if deleteData.TransactionID == id {
		t.Errorf("transaction id should be %s, but %s", id, deleteData.TransactionID)
		return
	}
}
