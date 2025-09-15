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
	CreateTr(ws *workspace.Workspace, record *models.TransactionRecord) error
	ListAllTrByLedgerId(ws *workspace.Workspace, ledgerId string) ([]*models.TransactionRecord, error)
	QueryTrsByPage(ws *workspace.Workspace, ledgerId string, offset, limit int) ([]*models.TransactionRecord, error)
	QueryCountOnCondition(ws *workspace.Workspace, ledgerId string) (int64, error)
	DeleteTrById(ws *workspace.Workspace, trId string) error
	CountTrByLedgerId(ws *workspace.Workspace, ledgerId string) (int64, error)
	DeleteAllTrByLedgerId(ws *workspace.Workspace, ledgerId string) error
}

var _ TransactionRecordDao = &transactionRecordDaoImpl{}

type transactionRecordDaoImpl struct{}

func (t *transactionRecordDaoImpl) CreateTr(ws *workspace.Workspace, record *models.TransactionRecord) error {
	if err := ws.GetDb().Create(record).Error; err != nil {
		return err
	}

	return nil
}

func (t *transactionRecordDaoImpl) ListAllTrByLedgerId(ws *workspace.Workspace, ledgerId string) ([]*models.TransactionRecord, error) {
	trs := make([]*models.TransactionRecord, 0)
	if err := ws.GetDb().
		Where("ledger_id = ?", ledgerId).
		Order("transaction_at desc, category desc").
		Find(&trs).Error; err != nil {
		return nil, err
	}

	return trs, nil
}

func (t *transactionRecordDaoImpl) QueryTrsByPage(ws *workspace.Workspace, ledgerId string, offset, limit int) ([]*models.TransactionRecord, error) {
	trs := make([]*models.TransactionRecord, 0)
	if err := ws.GetDb().
		Where("ledger_id = ?", ledgerId).
		Order("transaction_at desc, category desc").
		Offset(offset).
		Limit(limit).
		Find(&trs).Error; err != nil {
		return nil, err
	}

	return trs, nil
}

func (t *transactionRecordDaoImpl) QueryCountOnCondition(ws *workspace.Workspace, ledgerId string) (int64, error) {
	var count int64
	if err := ws.GetDb().Model(&models.TransactionRecord{}).
		Where("ledger_id = ?", ledgerId).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (t *transactionRecordDaoImpl) DeleteTrById(ws *workspace.Workspace, trId string) error {
	if err := ws.GetDb().
		Where("transaction_id = ?", trId).
		Delete(&models.TransactionRecord{}).Error; err != nil {
		return err
	}
	return nil
}

func (t *transactionRecordDaoImpl) CountTrByLedgerId(ws *workspace.Workspace, ledgerId string) (int64, error) {
	var count int64
	err := ws.GetDb().Model(&models.TransactionRecord{}).Where("ledger_id = ?", ledgerId).Count(&count).Error
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (t *transactionRecordDaoImpl) DeleteAllTrByLedgerId(ws *workspace.Workspace, ledgerId string) error {
	if err := ws.GetDb().Where("ledger_id = ?", ledgerId).Delete(&models.TransactionRecord{}).Error; err != nil {
		return err
	}
	return nil
}
