package dao

import (
	"sync"

	"github.com/billadm/models"
	"github.com/billadm/workspace"
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
		ledgerDao = &ledgerDaoImpl{}
	})
	return ledgerDao
}

type LedgerDao interface {
	CreateLedger(workspace *workspace.Workspace, ledger *models.Ledger) error
	ListAllLedger(workspace *workspace.Workspace) ([]models.Ledger, error)
	QueryLedgerById(workspace *workspace.Workspace, ledgerId string) (*models.Ledger, error)
	DeleteLedgerById(workspace *workspace.Workspace, ledgerId string) error
}

var _ LedgerDao = &ledgerDaoImpl{}

type ledgerDaoImpl struct{}

func (l *ledgerDaoImpl) CreateLedger(workspace *workspace.Workspace, ledger *models.Ledger) error {
	if err := workspace.GetDb().Create(ledger).Error; err != nil {
		return err
	}

	return nil
}

func (l *ledgerDaoImpl) ListAllLedger(workspace *workspace.Workspace) ([]models.Ledger, error) {
	ledgers := make([]models.Ledger, 0)
	if err := workspace.GetDb().Find(&ledgers).Error; err != nil {
		return nil, err
	}

	return ledgers, nil
}

func (l *ledgerDaoImpl) QueryLedgerById(workspace *workspace.Workspace, ledgerId string) (*models.Ledger, error) {
	var ledger models.Ledger
	if err := workspace.GetDb().Where("id = ?", ledgerId).First(&ledger).Error; err != nil {
		return nil, err
	}

	return &ledger, nil
}

func (l *ledgerDaoImpl) DeleteLedgerById(workspace *workspace.Workspace, ledgerId string) error {
	if err := workspace.GetDb().Where("id = ?", ledgerId).Delete(&models.Ledger{}).Error; err != nil {
		return err
	}

	return nil
}
