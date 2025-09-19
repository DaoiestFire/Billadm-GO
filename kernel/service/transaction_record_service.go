package service

import (
	"sync"

	"github.com/billadm/dao"
	"github.com/billadm/models"
	"github.com/billadm/models/dto"
	"github.com/billadm/util"
	"github.com/billadm/workspace"
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
	CreateTr(ws *workspace.Workspace, dto *dto.TransactionRecordDto) (string, error)
	ListAllTrsByLedgerId(ws *workspace.Workspace, ledgerId string) ([]*dto.TransactionRecordDto, error)
	QueryTrsOnCondition(ws *workspace.Workspace, condition *dto.QueryCondition) ([]*dto.TransactionRecordDto, error)
	QueryTrCountOnCondition(ws *workspace.Workspace, condition *dto.QueryCondition) (int64, error)
	DeleteTrById(ws *workspace.Workspace, trId string) error
}

var _ TransactionRecordService = &transactionRecordServiceImpl{}

type transactionRecordServiceImpl struct {
	trDao    dao.TransactionRecordDao
	trTagDao dao.TrTagDao
}

// CreateTr 创建成功返回交易记录的id
func (t *transactionRecordServiceImpl) CreateTr(ws *workspace.Workspace, trDto *dto.TransactionRecordDto) (string, error) {
	ws.GetLogger().Infof("start to create transaction record, ledger id: %s, description: %s", trDto.LedgerID, trDto.Description)

	transactionID := util.GetUUID()

	// 先创建消费记录
	record := trDto.ToTransactionRecord()
	record.TransactionID = transactionID
	if err := t.trDao.CreateTr(ws, record); err != nil {
		ws.GetLogger().Errorf("create transaction record failed, ledger id: %s, description: %s, err: %v", record.LedgerID, record.Description, err)
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
	if err := t.trTagDao.CreateTrTags(ws, trTags); err != nil {
		ws.GetLogger().Errorf("create trTags failed, ledger id: %s, description: %s, err: %v", record.LedgerID, record.Description, err)
		return "", err
	}

	ws.GetLogger().Infof("create transaction record success, ledger id: %s, description: %s", trDto.LedgerID, trDto.Description)
	return transactionID, nil
}

func (t *transactionRecordServiceImpl) ListAllTrsByLedgerId(ws *workspace.Workspace, ledgerId string) ([]*dto.TransactionRecordDto, error) {
	ws.GetLogger().Infof("start to list all transaction record, ledger id: %s", ledgerId)

	var err error
	// 先查询到所有的tr
	trs, err := t.trDao.ListAllTrByLedgerId(ws, ledgerId)
	if err != nil {
		return nil, err
	}

	// 再查询tr的tags进行组装
	trDtos := make([]*dto.TransactionRecordDto, 0, len(trs))
	for _, tr := range trs {
		trTags, err := t.trTagDao.QueryTrTagsByTrId(ws, tr.TransactionID)
		if err != nil {
			return nil, err
		}
		trDto := &dto.TransactionRecordDto{}
		trDto.FromTransactionRecord(tr)
		for _, tag := range trTags {
			trDto.Tags = append(trDto.Tags, tag.Tag)
		}
		trDtos = append(trDtos, trDto)
	}

	ws.GetLogger().Infof("list all transaction record success, ledger id: %s, len: %d", ledgerId, len(trs))
	return trDtos, err
}

func (t *transactionRecordServiceImpl) QueryTrsOnCondition(ws *workspace.Workspace, condition *dto.QueryCondition) ([]*dto.TransactionRecordDto, error) {
	ws.GetLogger().Infof("start to query trs, condition: %#v", condition)

	var err error
	// 先查询到所有的tr
	trs, err := t.trDao.QueryTrsByPage(ws, condition)
	if err != nil {
		return nil, err
	}

	// 再查询tr的tags进行组装
	trDtos := make([]*dto.TransactionRecordDto, 0, len(trs))
	for _, tr := range trs {
		trTags, err := t.trTagDao.QueryTrTagsByTrId(ws, tr.TransactionID)
		if err != nil {
			return nil, err
		}
		trDto := &dto.TransactionRecordDto{}
		trDto.FromTransactionRecord(tr)
		for _, tag := range trTags {
			trDto.Tags = append(trDto.Tags, tag.Tag)
		}
		trDtos = append(trDtos, trDto)
	}

	ws.GetLogger().Infof("query trs by page success, len: %d", len(trs))
	return trDtos, err
}

func (t *transactionRecordServiceImpl) QueryTrCountOnCondition(ws *workspace.Workspace, condition *dto.QueryCondition) (int64, error) {
	ws.GetLogger().Infof("start to query count of transaction records")
	cnt, err := t.trDao.QueryCountOnCondition(ws, condition)
	if err != nil {
		return 0, err
	}

	ws.GetLogger().Infof("query count of transaction records success, len: %d", cnt)
	return cnt, nil
}

func (t *transactionRecordServiceImpl) DeleteTrById(ws *workspace.Workspace, trId string) error {
	ws.GetLogger().Infof("start to delete transaction record, tr id: %s", trId)

	// 先删除消费记录的tags
	if err := t.trTagDao.DeleteTrTagByTrId(ws, trId); err != nil {
		return err
	}

	// 再删除对应的消费记录
	if err := t.trDao.DeleteTrById(ws, trId); err != nil {
		return err
	}

	ws.GetLogger().Infof("delete transaction record success, tr id: %s", trId)
	return nil
}
