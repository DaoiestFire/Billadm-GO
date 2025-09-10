package dao

import (
	"sync"

	"github.com/billadm/models"
	"github.com/billadm/workspace"
)

var (
	trTagDao     TrTagDao
	trTagDaoOnce sync.Once
)

func GetTrTagDao() TrTagDao {
	if trTagDao != nil {
		return trTagDao
	}

	trTagDaoOnce.Do(func() {
		trTagDao = &trTagDaoImpl{}
	})

	return trTagDao
}

type TrTagDao interface {
	CreateTrTags(workspace *workspace.Workspace, tags []*models.TrTag) error
	DeleteTrTagByLedgerId(workspace *workspace.Workspace, ledgerId string) error
	DeleteTrTagByTrId(workspace *workspace.Workspace, trId string) error
	QueryTrTagsByTrId(workspace *workspace.Workspace, trId string) ([]*models.TrTag, error)
}

var _ TrTagDao = &trTagDaoImpl{}

type trTagDaoImpl struct{}

func (t *trTagDaoImpl) CreateTrTags(workspace *workspace.Workspace, tags []*models.TrTag) error {
	if len(tags) <= 0 {
		return nil
	}
	if err := workspace.GetDb().Create(tags).Error; err != nil {
		return err
	}
	return nil
}

func (t *trTagDaoImpl) DeleteTrTagByLedgerId(workspace *workspace.Workspace, ledgerId string) error {
	if err := workspace.GetDb().Delete(&models.TrTag{}, "ledger_id = ?", ledgerId).Error; err != nil {
		return err
	}
	return nil
}

func (t *trTagDaoImpl) DeleteTrTagByTrId(workspace *workspace.Workspace, trId string) error {
	if err := workspace.GetDb().Delete(&models.TrTag{}, "transaction_id = ?", trId).Error; err != nil {
		return err
	}
	return nil
}

func (t *trTagDaoImpl) QueryTrTagsByTrId(workspace *workspace.Workspace, trId string) ([]*models.TrTag, error) {
	trTags := make([]*models.TrTag, 0)
	if err := workspace.GetDb().Where("transaction_id = ?", trId).Find(&trTags).Error; err != nil {
		return nil, err
	}
	return trTags, nil
}
