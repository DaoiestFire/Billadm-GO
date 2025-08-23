package service

import (
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/billadm/dao"
	"github.com/billadm/models"
)

var (
	tagService     TagService
	tagServiceOnce sync.Once
)

func GetTagService() TagService {
	if tagService != nil {
		return tagService
	}

	tagServiceOnce.Do(func() {
		tagService = &tagServiceImpl{
			tagDao: dao.GetTagDao(),
		}
	})

	return tagService
}

type TagService interface {
	QueryAllTag() ([]models.Tag, error)
}

var _ TagService = &tagServiceImpl{}

type tagServiceImpl struct {
	tagDao dao.TagDao
}

func (t *tagServiceImpl) QueryAllTag() ([]models.Tag, error) {
	logrus.Info("start to query all tag")
	tags, err := t.tagDao.QueryAllTags()
	if err != nil {
		return nil, err
	}

	logrus.Info("query all tag success, length: ", len(tags))
	return tags, nil
}
