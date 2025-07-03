package service

import (
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/billadm/kernel/dao"
	"github.com/billadm/kernel/models"
	"github.com/billadm/kernel/util"
)

var (
	trService     TransactionRecordService
	trServiceOnce sync.Once
)

func GetTrService() TransactionRecordService {
	if trService != nil {
		return trService
	}

	trServiceOnce.Do(func() {
		trService = &transactionRecordServiceImpl{
			trDao: dao.GetTrDao(),
		}
	})

	return trService
}

type TransactionRecordService interface {
	CreateTr(*models.TransactionRecord) (string, error)
	ListAllTrByLedgerId(ledgerId string) ([]*models.TransactionRecord, error)
}

var _ TransactionRecordService = &transactionRecordServiceImpl{}

type transactionRecordServiceImpl struct {
	trDao dao.TransactionRecordDao
}

// CreateTr 创建成功返回交易记录的id
func (t *transactionRecordServiceImpl) CreateTr(record *models.TransactionRecord) (string, error) {
	logrus.Infof("start to create transaction record, ledger id: %s, description: %s", record.LedgerID, record.Description)

	record.TransactionID = util.GetUUID()

	if err := t.trDao.CreateTr(record); err != nil {
		logrus.Errorf("create transaction record failed, ledger id: %s, description: %s, err: %v", record.LedgerID, record.Description, err)
		return "", err
	}

	logrus.Infof("create transaction record success, ledger id: %s, description: %s", record.LedgerID, record.Description)
	return record.TransactionID, nil
}

func (t *transactionRecordServiceImpl) ListAllTrByLedgerId(ledgerId string) ([]*models.TransactionRecord, error) {
	logrus.Infof("start to list all transaction record, ledger id: %s", ledgerId)

	trs, err := t.trDao.ListAllTrByLedgerId(ledgerId)
	if err != nil {
		return nil, err
	}

	logrus.Infof("list all transaction record success, ledger id: %s, len: %d", ledgerId, len(trs))
	return trs, err
}
