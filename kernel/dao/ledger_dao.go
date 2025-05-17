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
		ledgerDao = &LedgerDaoImpl{
			db: db.GetInstance(),
		}
	})
	return ledgerDao
}

type LedgerDao interface {
	CreateLedger(ledger *models.Ledger) error
	ListAllLedger(userId string) ([]models.Ledger, error)
	QueryLedgerById(ledgerId string) (*models.Ledger, error)
}

type LedgerDaoImpl struct {
	db *gorm.DB
}

var _ LedgerDao = &LedgerDaoImpl{}

func (l *LedgerDaoImpl) CreateLedger(ledger *models.Ledger) error {
	if err := l.db.Create(ledger).Error; err != nil {
		return err
	}

	return nil
}

func (l *LedgerDaoImpl) ListAllLedger(userId string) ([]models.Ledger, error) {
	ledgers := make([]models.Ledger, 0)
	if err := l.db.Find(&ledgers, "user_id", userId).Error; err != nil {
		return nil, err
	}

	return ledgers, nil
}

func (l *LedgerDaoImpl) QueryLedgerById(ledgerId string) (*models.Ledger, error) {
	var ledger models.Ledger
	if err := l.db.First(&ledger, ledgerId).Error; err != nil {
		return nil, err
	}

	return &ledger, nil
}
