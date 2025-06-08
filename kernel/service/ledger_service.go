package service

import (
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/billadm/kernel/dao"
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
		ledgerService = &ledgerServiceImpl{
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

var _ LedgerService = &ledgerServiceImpl{}

type ledgerServiceImpl struct {
	ledgerDao dao.LedgerDao
}

// CreateLedger 创建成功返回创建账本id
func (l *ledgerServiceImpl) CreateLedger(ledgerName, userId string) (string, error) {
	logrus.Infof("start to create ledger, name: %s, user id: %s", ledgerName, userId)
	// 账本名称可以重复，不需要校验账本名称
	ledger := &models.Ledger{
		ID:     util.GetUUID(),
		UserID: userId,
		Name:   ledgerName,
	}

	if err := l.ledgerDao.CreateLedger(ledger); err != nil {
		logrus.Errorf("create ledger failed, name: %s, user id: %s, err: %v", ledgerName, userId, err)
		return "", err
	}

	logrus.Infof("create ledger success, name: %s, user id: %s", ledgerName, userId)
	return ledger.ID, nil
}

// ListAllLedger 查询用户的所有账本
func (l *ledgerServiceImpl) ListAllLedger(userId string) ([]models.Ledger, error) {
	logrus.Infof("start to list all ledgers, user id: %s", userId)

	ledgers, err := l.ledgerDao.ListAllLedger(userId)
	if err != nil {
		logrus.Errorf("list all ledgers failed, user id: %s, err: %v", userId, err)
		return nil, err
	}

	logrus.Infof("end to list all ledgers, user id: %s, len: %d", userId, len(ledgers))
	return ledgers, nil
}

// QueryLedgerById 查询单个账本
func (l *ledgerServiceImpl) QueryLedgerById(ledgerId string) (*models.Ledger, error) {
	logrus.Infof("start to query ledger by id, id: %s", ledgerId)

	ledgers, err := l.ledgerDao.QueryLedgerById(ledgerId)
	if err != nil {
		logrus.Errorf("query ledger by id failed, id: %s, err: %v", ledgerId, err)
		return nil, err
	}

	logrus.Infof("end to query ledger by id, id: %s", ledgerId)
	return ledgers, nil
}
