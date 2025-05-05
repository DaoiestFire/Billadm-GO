package dao

import (
	"gorm.io/gorm"

	"github.com/billadm/kernel/models"
)

func GetLedgerDao() LedgerDao {

}

type LedgerDao interface {
	CreateLedger(ledger *models.Ledger) error
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
