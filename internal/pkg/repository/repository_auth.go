package repository

import (
	"context"
	"tugas_akhir_example/internal/daos"

	"gorm.io/gorm"
)

type AuthRepository interface {
	LoginUser(ctx context.Context, params daos.User) (res daos.User, err error)
	RegisterNewUser(ctx context.Context, data daos.User) (res daos.User, err error)
}

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{
		db: db,
	}
}

func (a *AuthRepositoryImpl) LoginUser(ctx context.Context, params daos.User) (res daos.User, err error) {
	db := a.db

	if err := db.Debug().WithContext(ctx).Where("notelp = ?", params.Notelp).First(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (a *AuthRepositoryImpl) RegisterNewUser(ctx context.Context, data daos.User) (res daos.User, err error) {
	result := a.db.WithContext(ctx).Create(&data)

	if result.Error != nil {
		return res, result.Error
	}

	return data, nil
}
