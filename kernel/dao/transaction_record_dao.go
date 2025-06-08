package dao

import (
	"sync"

	"gorm.io/gorm"

	"github.com/billadm/kernel/models"
	"github.com/billadm/kernel/util/db"
)

var (
	trDao     TransactionRecordDao
	trDaoOnce sync.Once
)

func GetTrDao() TransactionRecordDao {
	if trDao != nil {
		return trDao
	}

	trDaoOnce.Do(func() {
		trDao = &transactionRecordDaoImpl{
			db: db.GetInstance(),
		}
	})

	return trDao
}

type TransactionRecordDao interface {
	CreateTr(*models.TransactionRecord) error
}

var _ TransactionRecordDao = &transactionRecordDaoImpl{}

type transactionRecordDaoImpl struct {
	db *gorm.DB
}

func (t transactionRecordDaoImpl) CreateTr(record *models.TransactionRecord) error {
	//TODO implement me
	return nil
}
