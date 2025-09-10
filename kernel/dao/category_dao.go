package dao

import (
	"sync"

	"github.com/billadm/models"
	"github.com/billadm/workspace"
)

var (
	categoryDao     CategoryDao
	categoryDaoOnce sync.Once
)

func GetCategoryDao() CategoryDao {
	if categoryDao != nil {
		return categoryDao
	}
	categoryDaoOnce.Do(func() {
		categoryDao = &categoryDaoImpl{}
	})
	return categoryDao
}

type CategoryDao interface {
	QueryAllCategory(workspace *workspace.Workspace) ([]models.Category, error)
}

var _ CategoryDao = &categoryDaoImpl{}

type categoryDaoImpl struct{}

func (c *categoryDaoImpl) QueryAllCategory(workspace *workspace.Workspace) ([]models.Category, error) {
	categories := make([]models.Category, 0)
	if err := workspace.GetDb().Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
