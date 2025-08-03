package dao

import (
	"sync"

	"gorm.io/gorm"

	"github.com/billadm/models"
	"github.com/billadm/util/db"
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
	CreateTr(record *models.TransactionRecord) error
	ListAllTrByLedgerId(ledgerId string) ([]*models.TransactionRecord, error)
	DeleteTrById(trId string) error
	CountTrByLedgerId(ledgerId string) (int64, error)
	DeleteAllTrByLedgerId(ledgerId string) error
}

var _ TransactionRecordDao = &transactionRecordDaoImpl{}

type transactionRecordDaoImpl struct {
	db *gorm.DB
}

func (t *transactionRecordDaoImpl) CreateTr(record *models.TransactionRecord) error {
	if err := t.db.Create(record).Error; err != nil {
		return err
	}

	return nil
}

func (t *transactionRecordDaoImpl) ListAllTrByLedgerId(ledgerId string) ([]*models.TransactionRecord, error) {
	trs := make([]*models.TransactionRecord, 0)
	if err := t.db.Where("ledger_id = ?", ledgerId).Find(&trs).Error; err != nil {
		return nil, err
	}

	return trs, nil
}

func (t *transactionRecordDaoImpl) DeleteTrById(trId string) error {
	if err := t.db.Where("transaction_id = ?", trId).Delete(&models.TransactionRecord{}).Error; err != nil {
		return err
	}
	return nil
}

func (t *transactionRecordDaoImpl) CountTrByLedgerId(ledgerId string) (int64, error) {
	var count int64
	err := t.db.Model(&models.TransactionRecord{}).Where("ledger_id = ?", ledgerId).Count(&count).Error
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (t *transactionRecordDaoImpl) DeleteAllTrByLedgerId(ledgerId string) error {
	if err := t.db.Where("ledger_id = ?", ledgerId).Delete(&models.TransactionRecord{}).Error; err != nil {
		return err
	}
	return nil
}
