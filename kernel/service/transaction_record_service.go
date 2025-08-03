package service

import (
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/billadm/dao"
	"github.com/billadm/models"
	"github.com/billadm/models/dto"
	"github.com/billadm/util"
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
			trDao:    dao.GetTrDao(),
			trTagDao: dao.GetTrTagDao(),
		}
	})

	return trService
}

type TransactionRecordService interface {
	CreateTr(dto *dto.TransactionRecordDto) (string, error)
	ListAllTrByLedgerId(ledgerId string) ([]*models.TransactionRecord, error)
	DeleteTrById(string) error
}

var _ TransactionRecordService = &transactionRecordServiceImpl{}

type transactionRecordServiceImpl struct {
	trDao    dao.TransactionRecordDao
	trTagDao dao.TrTagDao
}

// CreateTr 创建成功返回交易记录的id
func (t *transactionRecordServiceImpl) CreateTr(trDto *dto.TransactionRecordDto) (string, error) {
	logrus.Infof("start to create transaction record, ledger id: %s, description: %s", trDto.LedgerID, trDto.Description)

	transactionID := util.GetUUID()

	// 先创建消费记录
	record := trDto.ToTransactionRecord()
	record.TransactionID = transactionID
	if err := t.trDao.CreateTr(record); err != nil {
		logrus.Errorf("create transaction record failed, ledger id: %s, description: %s, err: %v", record.LedgerID, record.Description, err)
		return "", err
	}

	// 再插入消费记录的tag
	trTags := make([]*models.TrTag, 0, len(trDto.Tags))
	for _, tag := range trDto.Tags {
		trTag := &models.TrTag{
			LedgerID:      trDto.LedgerID,
			TransactionID: transactionID,
			Tag:           tag,
		}
		trTags = append(trTags, trTag)
	}
	if err := t.trTagDao.CreateTrTags(trTags); err != nil {
		logrus.Errorf("create trTags failed, ledger id: %s, description: %s, err: %v", record.LedgerID, record.Description, err)
		return "", err
	}

	logrus.Infof("create transaction record success, ledger id: %s, description: %s", trDto.LedgerID, trDto.Description)
	return transactionID, nil
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

func (t *transactionRecordServiceImpl) DeleteTrById(trId string) error {
	logrus.Infof("start to delete transaction record, tr id: %s", trId)

	// 先删除消费记录的tags
	if err := t.trTagDao.DeleteTrTagByTrId(trId); err != nil {
		return err
	}

	// 再删除对应的消费记录
	if err := t.trDao.DeleteTrById(trId); err != nil {
		return err
	}

	logrus.Infof("delete transaction record success, tr id: %s", trId)
	return nil
}
