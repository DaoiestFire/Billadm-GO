package dao

import (
	"sync"

	"gorm.io/gorm"

	"github.com/billadm/models"
	"github.com/billadm/util/db"
)

var (
	tagDao     TagDao
	tagDaoOnce sync.Once
)

func GetTagDao() TagDao {
	if tagDao != nil {
		return tagDao
	}
	tagDaoOnce.Do(func() {
		tagDao = &TagDaoImpl{
			db: db.GetInstance(),
		}
	})
	return tagDao
}

type TagDao interface {
	QueryAllTags() ([]models.Tag, error)
}

var _ TagDao = &TagDaoImpl{}

type TagDaoImpl struct {
	db *gorm.DB
}

func (t *TagDaoImpl) QueryAllTags() ([]models.Tag, error) {
	tags := make([]models.Tag, 0)
	if err := t.db.Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}
