package dao

import (
	"gorm.io/gorm"
	"sync"

	"github.com/billadm/util/db"

	"github.com/billadm/models"
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
		trTagDao = &trTagDaoImpl{
			db: db.GetInstance(),
		}
	})

	return trTagDao
}

type TrTagDao interface {
	CreateTrTags([]*models.TrTag) error
	DeleteTrTagByLedgerId(ledgerId string) error
	DeleteTrTagByTrId(trId string) error
	QueryTrTagsByTrId(trId string) ([]*models.TrTag, error)
}

var _ TrTagDao = &trTagDaoImpl{}

type trTagDaoImpl struct {
	db *gorm.DB
}

func (t *trTagDaoImpl) CreateTrTags(tags []*models.TrTag) error {
	//TODO implement me
	panic("implement me")
}

func (t *trTagDaoImpl) DeleteTrTagByLedgerId(ledgerId string) error {
	//TODO implement me
	panic("implement me")
}

func (t *trTagDaoImpl) DeleteTrTagByTrId(trId string) error {
	//TODO implement me
	panic("implement me")
}

func (t *trTagDaoImpl) QueryTrTagsByTrId(trId string) ([]*models.TrTag, error) {
	//TODO implement me
	panic("implement me")
}
