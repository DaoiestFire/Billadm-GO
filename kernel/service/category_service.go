package service

import (
	"sync"

	"github.com/billadm/dao"
	"github.com/billadm/models"
	"github.com/sirupsen/logrus"
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
	QueryAllCategory() ([]models.Category, error)
}

var _ CategoryService = &categoryServiceImpl{}

type categoryServiceImpl struct {
	categoryDao dao.CategoryDao
}

func (c *categoryServiceImpl) QueryAllCategory() ([]models.Category, error) {
	logrus.Info("start to query all category")
	categories, err := c.categoryDao.QueryAllCategory()
	if err != nil {
		return nil, err
	}

	logrus.Info("query all category success, length: ", len(categories))
	return categories, nil
}
