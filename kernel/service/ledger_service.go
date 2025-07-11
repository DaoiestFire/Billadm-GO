package service

import (
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/billadm/dao"
	"github.com/billadm/models"
	"github.com/billadm/util"
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
			trDao:     dao.GetTrDao(),
		}
	})

	return ledgerService
}

type LedgerService interface {
	CreateLedger(ledgerName string) (string, error)
	ListAllLedger() ([]models.Ledger, error)
	QueryLedgerById(ledgerId string) (*models.Ledger, error)
	DeleteLedgerById(ledgerId string) error
}

var _ LedgerService = &ledgerServiceImpl{}

type ledgerServiceImpl struct {
	ledgerDao dao.LedgerDao
	trDao     dao.TransactionRecordDao
}

// CreateLedger 创建成功返回创建账本id
func (l *ledgerServiceImpl) CreateLedger(ledgerName string) (string, error) {
	logrus.Infof("start to create ledger, name: %s", ledgerName)
	// 账本名称可以重复，不需要校验账本名称
	ledger := &models.Ledger{
		ID:   util.GetUUID(),
		Name: ledgerName,
	}

	if err := l.ledgerDao.CreateLedger(ledger); err != nil {
		logrus.Errorf("create ledger failed, name: %s, err: %v", ledgerName, err)
		return "", err
	}

	logrus.Infof("create ledger success, name: %s", ledgerName)
	return ledger.ID, nil
}

// ListAllLedger 查询所有账本
func (l *ledgerServiceImpl) ListAllLedger() ([]models.Ledger, error) {
	logrus.Infof("start to list all ledgers")

	ledgers, err := l.ledgerDao.ListAllLedger()
	if err != nil {
		logrus.Errorf("list all ledgers failed, err: %v", err)
		return nil, err
	}

	logrus.Infof("end to list all ledgers, len: %d", len(ledgers))
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

func (l *ledgerServiceImpl) DeleteLedgerById(ledgerId string) error {
	logrus.Infof("start to delete ledger by id, id: %s", ledgerId)

	if err := l.ledgerDao.DeleteLedgerById(ledgerId); err != nil {
		logrus.Errorf("delete ledger by id failed, id: %s, err: %v", ledgerId, err)
		return err
	}

	cnt, err := l.trDao.CountTrByLedgerId(ledgerId)
	if err != nil {
		logrus.Errorf("get count of trs from ledger by id failed, id: %s, err: %v", ledgerId, err)
		return err
	}
	logrus.Infof("will delete trs by ledger id: %s, count: %d", ledgerId, cnt)

	if err := l.trDao.DeleteAllTrByLedgerId(ledgerId); err != nil {
		logrus.Errorf("delete all tr by ledger id failed, id: %s, err: %v", ledgerId, err)
		return err
	}

	logrus.Infof("end to delete ledger by id, id: %s", ledgerId)
	return nil
}
