package repository

import (
	"context"
	"tugas_akhir_example/internal/daos"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCategory(ctx context.Context) (res []daos.Category, err error)
	GetCategoryByID(ctx context.Context, categoryid string) (res daos.Category, err error)
	CreateCategory(ctx context.Context, data daos.Category) (res uint, err error)
	UpdateCategoryByID(ctx context.Context, categoryid string, data daos.Category) (res string, err error)
	DeleteCategoryByID(ctx context.Context, categoryid string) (res string, err error)
}

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		db: db,
	}
}
func (a *CategoryRepositoryImpl) GetAllCategory(ctx context.Context) (res []daos.Category, err error) {
	db := a.db

	if err := db.Debug().WithContext(ctx).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (a *CategoryRepositoryImpl) GetCategoryByID(ctx context.Context, categoryid string) (res daos.Category, err error) {
	if err := a.db.First(&res, categoryid).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (a *CategoryRepositoryImpl) CreateCategory(ctx context.Context, data daos.Category) (res uint, err error) {
	result := a.db.Create(&data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}

	return data.Id, nil
}

func (a *CategoryRepositoryImpl) UpdateCategoryByID(ctx context.Context, categoryid string, data daos.Category) (res string, err error) {
	var dataCategory daos.Category
	if err = a.db.Where("id = ? ", categoryid).First(&dataCategory).WithContext(ctx).Error; err != nil {
		return "Update Category failed", gorm.ErrRecordNotFound
	}

	if err := a.db.Model(dataCategory).Updates(&data).Where("id = ? ", categoryid).Error; err != nil {
		return "Update Category failed", err
	}

	return res, nil
}

func (a *CategoryRepositoryImpl) DeleteCategoryByID(ctx context.Context, categoryid string) (res string, err error) {
	var dataCategory daos.Category
	if err = a.db.Where("id = ?", categoryid).First(&dataCategory).WithContext(ctx).Error; err != nil {
		return "Delete book failed", gorm.ErrRecordNotFound
	}

	if err := a.db.Model(dataCategory).Delete(&dataCategory).Error; err != nil {
		return "Delete book failed", err
	}

	return res, nil
}
