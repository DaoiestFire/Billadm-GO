package service

import (
	"sync"

	"github.com/billadm/dao"
	"github.com/billadm/models"
	"github.com/billadm/workspace"
)

var (
	categoryService     CategoryService
	categoryServiceOnce sync.Once
)

func GetCategoryService() CategoryService {
	if categoryService != nil {
		return categoryService
	}

	categoryServiceOnce.Do(func() {
		categoryService = &categoryServiceImpl{
			categoryDao: dao.GetCategoryDao(),
		}
	})

	return categoryService
}

type CategoryService interface {
	QueryAllCategory(ws *workspace.Workspace) ([]models.Category, error)
}

var _ CategoryService = &categoryServiceImpl{}

type categoryServiceImpl struct {
	categoryDao dao.CategoryDao
}

func (c *categoryServiceImpl) QueryAllCategory(ws *workspace.Workspace) ([]models.Category, error) {
	ws.GetLogger().Info("start to query all category")
	categories, err := c.categoryDao.QueryAllCategory(ws)
	if err != nil {
		return nil, err
	}

	ws.GetLogger().Info("query all category success, length: ", len(categories))
	return categories, nil
}
