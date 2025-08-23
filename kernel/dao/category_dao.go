package dao

import (
	"gorm.io/gorm"
	"sync"

	"github.com/billadm/models"
	"github.com/billadm/util/db"
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
		categoryDao = &categoryDaoImpl{
			db: db.GetInstance(),
		}
	})
	return categoryDao
}

type CategoryDao interface {
	QueryAllCategory() ([]models.Category, error)
}

var _ CategoryDao = &categoryDaoImpl{}

type categoryDaoImpl struct {
	db *gorm.DB
}

func (c *categoryDaoImpl) QueryAllCategory() ([]models.Category, error) {
	categories := make([]models.Category, 0)
	if err := c.db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}
