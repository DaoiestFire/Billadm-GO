package dao

import (
	"gorm.io/gorm"
	"sync"

	"github.com/billadm/util/db"
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
	DeleteTrTagByTrId(trId string) error
}

var _ TrTagDao = &trTagDaoImpl{}

type trTagDaoImpl struct {
	db *gorm.DB
}

func (t *trTagDaoImpl) DeleteTrTagByTrId(trId string) error {
	//TODO implement me
	panic("implement me")
}
