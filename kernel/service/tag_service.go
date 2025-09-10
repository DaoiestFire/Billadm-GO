package service

import (
	"sync"

	"github.com/billadm/dao"
	"github.com/billadm/models"
	"github.com/billadm/workspace"
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
	QueryAllTag(ws *workspace.Workspace) ([]models.Tag, error)
}

var _ TagService = &tagServiceImpl{}

type tagServiceImpl struct {
	tagDao dao.TagDao
}

func (t *tagServiceImpl) QueryAllTag(ws *workspace.Workspace) ([]models.Tag, error) {
	ws.GetLogger().Info("start to query all tag")
	tags, err := t.tagDao.QueryAllTags(ws)
	if err != nil {
		return nil, err
	}

	ws.GetLogger().Info("query all tag success, length: ", len(tags))
	return tags, nil
}
