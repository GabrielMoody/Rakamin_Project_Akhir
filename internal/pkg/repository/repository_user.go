package repository

import (
	"context"
	"tugas_akhir_example/internal/daos"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetMyProfile(ctx context.Context, userid string) (res daos.User, err error)
	UpdateMyProfile(ctx context.Context, userid string, data daos.User) (res daos.User, err error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (a *UserRepositoryImpl) GetMyProfile(ctx context.Context, userid string) (res daos.User, err error) {
	if err := a.db.WithContext(ctx).First(&res, userid).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (a *UserRepositoryImpl) UpdateMyProfile(ctx context.Context, userid string, data daos.User) (res daos.User, err error) {
	if err := a.db.WithContext(ctx).First(&res, userid).Error; err != nil {
		return res, gorm.ErrRecordNotFound
	}

	if err := a.db.Model(&res).Updates(data).Error; err != nil {
		return res, err
	}

	return res, nil
}
