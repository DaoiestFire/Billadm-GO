package dao

import (
	"sync"

	"gorm.io/gorm"

	"github.com/billadm/models"
	"github.com/billadm/util/db"
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
	ListAllLedger() ([]models.Ledger, error)
	QueryLedgerById(ledgerId string) (*models.Ledger, error)
	DeleteLedgerById(ledgerId string) error
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

func (l *ledgerDaoImpl) ListAllLedger() ([]models.Ledger, error) {
	ledgers := make([]models.Ledger, 0)
	if err := l.db.Find(&ledgers).Error; err != nil {
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

func (l *ledgerDaoImpl) DeleteLedgerById(ledgerId string) error {
	if err := l.db.Where("id = ?", ledgerId).Delete(&models.Ledger{}).Error; err != nil {
		return err
	}

	return nil
}
