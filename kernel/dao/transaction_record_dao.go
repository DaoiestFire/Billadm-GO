package dao

import (
	"database/sql"
	"sync"

	"github.com/billadm/models"
	"github.com/billadm/models/dto"
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
	QueryTrsOnCondition(ws *workspace.Workspace, condition *dto.TrQueryCondition) ([]*models.TransactionRecord, error)
	QueryCountOnCondition(ws *workspace.Workspace, condition *dto.TrQueryCondition) (int64, error)
	QueryPriceOnCondition(ws *workspace.Workspace, condition *dto.TrQueryCondition) (float64, error)
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

func (t *transactionRecordDaoImpl) QueryTrsOnCondition(ws *workspace.Workspace, condition *dto.TrQueryCondition) ([]*models.TransactionRecord, error) {
	trs := make([]*models.TransactionRecord, 0)
	db := ws.GetDb().Where("ledger_id = ?", condition.LedgerID)
	db = db.Order("transaction_at desc, transaction_type asc, category desc, price desc")
	if condition.Offset != -1 && condition.Limit != -1 {
		db = db.Offset(condition.Offset).Limit(condition.Limit)
	}
	if len(condition.TsRange) == 2 {
		db = db.Where("transaction_at >= ?", condition.TsRange[0]).Where("transaction_at <= ?", condition.TsRange[1])
	}
	if len(condition.TransactionType) > 0 {
		db = db.Where("transaction_type IN (?)", condition.TransactionType)
	}
	db = db.Find(&trs)
	if err := db.Error; err != nil {
		return nil, err
	}
	return trs, nil
}

func (t *transactionRecordDaoImpl) QueryCountOnCondition(ws *workspace.Workspace, condition *dto.TrQueryCondition) (int64, error) {
	var count int64
	db := ws.GetDb().Model(&models.TransactionRecord{})
	db = db.Where("ledger_id = ?", condition.LedgerID)
	if len(condition.TsRange) == 2 {
		db = db.Where("transaction_at >= ?", condition.TsRange[0]).Where("transaction_at <= ?", condition.TsRange[1])
	}
	db = db.Count(&count)
	if err := db.Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (t *transactionRecordDaoImpl) QueryPriceOnCondition(ws *workspace.Workspace, condition *dto.TrQueryCondition) (float64, error) {
	var price sql.NullFloat64
	db := ws.GetDb().Model(&models.TransactionRecord{})
	db = db.Where("ledger_id = ?", condition.LedgerID)
	db = db.Order("transaction_at desc, transaction_type asc, category desc, price desc")
	if condition.Offset != -1 && condition.Limit != -1 {
		db = db.Offset(condition.Offset).Limit(condition.Limit)
	}
	if len(condition.TsRange) == 2 {
		db = db.Where("transaction_at >= ?", condition.TsRange[0]).Where("transaction_at <= ?", condition.TsRange[1])
	}
	if len(condition.TransactionType) > 0 {
		db = db.Where("transaction_type IN (?)", condition.TransactionType)
	}
	db = db.Select("SUM(price)").Scan(&price)
	if err := db.Error; err != nil {
		return 0, err
	}
	return price.Float64, nil
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
