package dao

import (
	"sync"

	"github.com/billadm/models"
	"github.com/billadm/workspace"
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
		trDao = &transactionRecordDaoImpl{}
	})

	return trDao
}

type TransactionRecordDao interface {
	CreateTr(workspace *workspace.Workspace, record *models.TransactionRecord) error
	ListAllTrByLedgerId(workspace *workspace.Workspace, ledgerId string) ([]*models.TransactionRecord, error)
	QueryTrsByPage(workspace *workspace.Workspace, ledgerId string, offset, limit int) ([]*models.TransactionRecord, error)
	QueryCountOnCondition(workspace *workspace.Workspace, ledgerId string) (int64, error)
	DeleteTrById(workspace *workspace.Workspace, trId string) error
	CountTrByLedgerId(workspace *workspace.Workspace, ledgerId string) (int64, error)
	DeleteAllTrByLedgerId(workspace *workspace.Workspace, ledgerId string) error
}

var _ TransactionRecordDao = &transactionRecordDaoImpl{}

type transactionRecordDaoImpl struct{}

func (t *transactionRecordDaoImpl) CreateTr(workspace *workspace.Workspace, record *models.TransactionRecord) error {
	if err := workspace.GetDb().Create(record).Error; err != nil {
		return err
	}

	return nil
}

func (t *transactionRecordDaoImpl) ListAllTrByLedgerId(workspace *workspace.Workspace, ledgerId string) ([]*models.TransactionRecord, error) {
	trs := make([]*models.TransactionRecord, 0)
	if err := workspace.GetDb().Where("ledger_id = ?", ledgerId).Find(&trs).Error; err != nil {
		return nil, err
	}

	return trs, nil
}

func (t *transactionRecordDaoImpl) QueryTrsByPage(workspace *workspace.Workspace, ledgerId string, offset, limit int) ([]*models.TransactionRecord, error) {
	trs := make([]*models.TransactionRecord, 0)
	if err := workspace.GetDb().Where("ledger_id = ?", ledgerId).Offset(offset).Limit(limit).Find(&trs).Error; err != nil {
		return nil, err
	}

	return trs, nil
}

func (t *transactionRecordDaoImpl) QueryCountOnCondition(workspace *workspace.Workspace, ledgerId string) (int64, error) {
	var count int64
	if err := workspace.GetDb().Model(&models.TransactionRecord{}).Where("ledger_id = ?", ledgerId).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (t *transactionRecordDaoImpl) DeleteTrById(workspace *workspace.Workspace, trId string) error {
	if err := workspace.GetDb().Where("transaction_id = ?", trId).Delete(&models.TransactionRecord{}).Error; err != nil {
		return err
	}
	return nil
}

func (t *transactionRecordDaoImpl) CountTrByLedgerId(workspace *workspace.Workspace, ledgerId string) (int64, error) {
	var count int64
	err := workspace.GetDb().Model(&models.TransactionRecord{}).Where("ledger_id = ?", ledgerId).Count(&count).Error
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (t *transactionRecordDaoImpl) DeleteAllTrByLedgerId(workspace *workspace.Workspace, ledgerId string) error {
	if err := workspace.GetDb().Where("ledger_id = ?", ledgerId).Delete(&models.TransactionRecord{}).Error; err != nil {
		return err
	}
	return nil
}
