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
	ListAllTrByLedgerId(string) ([]*models.TransactionRecord, error)
}

var _ TransactionRecordDao = &transactionRecordDaoImpl{}

type transactionRecordDaoImpl struct {
	db *gorm.DB
}

func (t transactionRecordDaoImpl) CreateTr(record *models.TransactionRecord) error {
	if err := t.db.Create(record).Error; err != nil {
		return err
	}

	return nil
}

func (t transactionRecordDaoImpl) ListAllTrByLedgerId(ledgerId string) ([]*models.TransactionRecord, error) {
	trs := make([]*models.TransactionRecord, 0)
	if err := t.db.Where("ledger_id = ?", ledgerId).Find(&trs).Error; err != nil {
		return nil, err
	}

	return trs, nil
}
