package repository

import (
	"context"
	"fmt"
	"tugas_akhir_example/internal/daos"

	"gorm.io/gorm"
)

type RepositoryToko interface {
	GetMyToko(ctx context.Context, userid string) (toko daos.Toko, err error)
	GetAllToko(ctx context.Context, params daos.FilterToko) (toko []daos.Toko, err error)
	GetTokoByID(ctx context.Context, tokoid string) (toko daos.Toko, err error)
	CreateTokoByUserRegistration(ctx context.Context, user daos.User) (res uint, err error)
	UpdateTokoByID(ctx context.Context, tokoid string, data daos.Toko) (res string, err error)
}

type RepositoryTokoImpl struct {
	db *gorm.DB
}

func NewRepositoryToko(db *gorm.DB) RepositoryToko {
	return &RepositoryTokoImpl{
		db: db,
	}
}

func (a *RepositoryTokoImpl) GetMyToko(ctx context.Context, userid string) (toko daos.Toko, err error) {
	if err := a.db.Debug().WithContext(ctx).First(&toko, "id_user = ?", userid).Error; err != nil {
		return toko, err
	}

	return toko, nil
}

func (a *RepositoryTokoImpl) GetAllToko(ctx context.Context, params daos.FilterToko) (toko []daos.Toko, err error) {
	db := a.db

	filter := map[string][]any{
		"nama like ?": {fmt.Sprint("%" + params.Name)},
	}

	for key, val := range filter {
		db = db.Where(key, val...)
	}

	if err := db.Debug().WithContext(ctx).Limit(params.Limit).Offset(params.Offset).Find(&toko).Error; err != nil {
		return toko, err
	}

	return toko, nil
}

func (a *RepositoryTokoImpl) GetTokoByID(ctx context.Context, tokoid string) (toko daos.Toko, err error) {
	if err := a.db.Debug().WithContext(ctx).First(&toko, tokoid).Error; err != nil {
		return toko, err
	}

	return toko, nil
}

func (a *RepositoryTokoImpl) CreateTokoByUserRegistration(ctx context.Context, user daos.User) (res uint, err error) {
	dataToko := daos.Toko{
		Id_user:   user.Id,
		Nama_toko: user.Nama,
	}

	result := a.db.Create(&dataToko).WithContext(ctx)

	if result.Error != nil {
		return res, result.Error
	}

	return dataToko.Id, nil
}

func (a *RepositoryTokoImpl) UpdateTokoByID(ctx context.Context, tokoid string, data daos.Toko) (res string, err error) {
	var dataToko daos.Toko

	if err := a.db.Where("id = ?", tokoid).First(&dataToko).WithContext(ctx).Error; err != nil {
		return "Data is not found", gorm.ErrRecordNotFound
	}

	if err := a.db.Model(dataToko).Updates(&data).Where("id = ?", tokoid).Error; err != nil {
		return "Error updating data", err
	}

	return res, err
}
