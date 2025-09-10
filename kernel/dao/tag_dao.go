package dao

import (
	"sync"

	"github.com/billadm/models"
	"github.com/billadm/workspace"
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
		tagDao = &TagDaoImpl{}
	})
	return tagDao
}

type TagDao interface {
	QueryAllTags(workspace *workspace.Workspace) ([]models.Tag, error)
}

var _ TagDao = &TagDaoImpl{}

type TagDaoImpl struct{}

func (t *TagDaoImpl) QueryAllTags(workspace *workspace.Workspace) ([]models.Tag, error) {
	tags := make([]models.Tag, 0)
	if err := workspace.GetDb().Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}
