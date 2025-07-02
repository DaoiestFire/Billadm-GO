package dao

import (
	"sync"

	"gorm.io/gorm"

	"github.com/billadm/kernel/models"
	"github.com/billadm/kernel/util/db"
)

var (
	ledgerDao     LedgerDao
	ledgerDaoOnce sync.Once
)

func GetLedgerDao() LedgerDao {
	if ledgerDao != nil {
		return ledgerDao
	}
	ledgerDaoOnce.Do(func() {
		ledgerDao = &ledgerDaoImpl{
			db: db.GetInstance(),
		}
	})
	return ledgerDao
}

type LedgerDao interface {
	CreateLedger(ledger *models.Ledger) error
	ListAllLedgerByUserId(userId string) ([]models.Ledger, error)
	QueryLedgerById(ledgerId string) (*models.Ledger, error)
}

var _ LedgerDao = &ledgerDaoImpl{}

type ledgerDaoImpl struct {
	db *gorm.DB
}

func (l *ledgerDaoImpl) CreateLedger(ledger *models.Ledger) error {
	if err := l.db.Create(ledger).Error; err != nil {
		return err
	}

	return nil
}

func (l *ledgerDaoImpl) ListAllLedgerByUserId(userId string) ([]models.Ledger, error) {
	ledgers := make([]models.Ledger, 0)
	if err := l.db.Where("user_id = ?", userId).Find(&ledgers).Error; err != nil {
		return nil, err
	}

	return ledgers, nil
}

func (l *ledgerDaoImpl) QueryLedgerById(ledgerId string) (*models.Ledger, error) {
	var ledger models.Ledger
	if err := l.db.Where("id = ?", ledgerId).First(&ledger).Error; err != nil {
		return nil, err
	}

	return &ledger, nil
}
