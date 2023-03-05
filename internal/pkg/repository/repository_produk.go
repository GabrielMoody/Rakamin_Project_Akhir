package repository

import (
	"context"
	"fmt"
	"tugas_akhir_example/internal/daos"

	"gorm.io/gorm"
)

type ProdukRepository interface {
	GetAllProduk(ctx context.Context, params daos.FilterProduk) (data []daos.Produk, err error)
	GetProdukByID(ctx context.Context, produkid string) (data daos.Produk, err error)
	CreateProduk(ctx context.Context, data daos.Produk) (res uint, err error)
	UpdateProdukByID(ctx context.Context, produkid string, data daos.Produk) (res string, err error)
	DeleteProdukByID(ctx context.Context, produkid string) (res string, err error)
}

type ProdukRepositoryImpl struct {
	db *gorm.DB
}

func NewProdukRepository(db *gorm.DB) ProdukRepository {
	return &ProdukRepositoryImpl{
		db: db,
	}
}
func (a *ProdukRepositoryImpl) GetAllProduk(ctx context.Context, params daos.FilterProduk) (data []daos.Produk, err error) {
	db := a.db

	filter := map[string][]any{
		"nama_produk like ?": {fmt.Sprint("%" + params.Nama_produk)},
	}

	for key, val := range filter {
		db = db.Where(key, val...)
	}

	if err := db.Debug().WithContext(ctx).Limit(params.Limit).Offset(params.Offset).Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (a *ProdukRepositoryImpl) GetProdukByID(ctx context.Context, produkid string) (data daos.Produk, err error) {
	if err := a.db.First(&data, produkid).WithContext(ctx).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (a *ProdukRepositoryImpl) CreateProduk(ctx context.Context, data daos.Produk) (res uint, err error) {
	result := a.db.Create(&data).WithContext(ctx)

	if result.Error != nil {
		return res, result.Error
	}

	return data.Id, nil
}

func (a *ProdukRepositoryImpl) UpdateProdukByID(ctx context.Context, produkid string, data daos.Produk) (res string, err error) {
	var dataProduk daos.Produk
	if err := a.db.Where("id = ?", produkid).First(&data).WithContext(ctx).Error; err != nil {
		return "Error update data", gorm.ErrRecordNotFound
	}

	if err := a.db.Model(dataProduk).Updates(&data).Where("id = ?", produkid).Error; err != nil {
		return "Error update data", err
	}

	return res, nil
}

func (a *ProdukRepositoryImpl) DeleteProdukByID(ctx context.Context, produkid string) (res string, err error) {
	var dataProduk daos.Produk
	if err := a.db.First(&dataProduk, produkid).WithContext(ctx).Error; err != nil {
		return "Error delete data", gorm.ErrRecordNotFound
	}

	if err := a.db.Model(dataProduk).Delete(&dataProduk).Error; err != nil {
		return "Error delete data", err
	}

	return res, nil
}
