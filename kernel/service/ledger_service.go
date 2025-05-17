package service

import (
	"sync"

	"github.com/billadm/kernel/dao"
	"github.com/billadm/kernel/logger"
	"github.com/billadm/kernel/models"
	"github.com/billadm/kernel/util"
)

var (
	ledgerService     LedgerService
	ledgerServiceOnce sync.Once
)

func GetLedgerService() LedgerService {
	if ledgerService != nil {
		return ledgerService
	}

	ledgerServiceOnce.Do(func() {
		ledgerService = &LedgerServiceImpl{
			ledgerDao: dao.GetLedgerDao(),
		}
	})

	return ledgerService
}

type LedgerService interface {
	CreateLedger(ledgerName, userId string) (string, error)
	ListAllLedger(userId string) ([]models.Ledger, error)
	QueryLedgerById(ledgerId string) (*models.Ledger, error)
}

type LedgerServiceImpl struct {
	ledgerDao dao.LedgerDao
}

var _ LedgerService = &LedgerServiceImpl{}

// CreateLedger 创建成功返回创建成功的账本id
func (l *LedgerServiceImpl) CreateLedger(ledgerName, userId string) (string, error) {
	logger.Info("start to create ledger, name: %s, user id: %s", ledgerName, userId)
	// 账本名称可以重复，不需要校验账本名称
	ledger := &models.Ledger{
		ID:     util.GetUUID(),
		UserID: userId,
		Name:   ledgerName,
	}

	if err := l.ledgerDao.CreateLedger(ledger); err != nil {
		logger.Error("create ledger failed, name: %s, user id: %s, err: %v", ledgerName, userId, err)
		return "", err
	}

	logger.Info("create ledger success, name: %s, user id: %s", ledgerName, userId)
	return ledger.ID, nil
}

// ListAllLedger 查询用户的所有账本
func (l *LedgerServiceImpl) ListAllLedger(userId string) ([]models.Ledger, error) {
	logger.Info("start to list all ledgers, user id: %s", userId)

	ledgers, err := l.ledgerDao.ListAllLedger(userId)
	if err != nil {
		logger.Error("list all ledgers failed, user id: %s, err: %v", userId, err)
		return nil, err
	}

	logger.Info("end to list all ledgers, user id: %s, len: %d", userId, len(ledgers))
	return ledgers, nil
}

// QueryLedgerById 查询单个账本
func (l *LedgerServiceImpl) QueryLedgerById(ledgerId string) (*models.Ledger, error) {
	logger.Info("start to query ledger by id, id: %s", ledgerId)

	ledgers, err := l.ledgerDao.QueryLedgerById(ledgerId)
	if err != nil {
		logger.Error("query ledger by id failed, id: %s, err: %v", ledgerId, err)
		return nil, err
	}

	logger.Info("end to query ledger by id, id: %s", ledgerId)
	return ledgers, nil
}
