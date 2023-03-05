package repository

import (
	"context"
	"strconv"
	"tugas_akhir_example/internal/daos"

	"gorm.io/gorm"
)

type AlamatRepository interface {
	GetMyAlamat(ctx context.Context, userid string) (res daos.Alamat, err error)
	GetAlamatByID(ctx context.Context, alamatid string) (res daos.Alamat, err error)
	CreateAlamat(ctx context.Context, userid string, data daos.Alamat) (res uint, err error)
	UpdateAlamatByID(ctx context.Context, alamatid string, data daos.Alamat) (res string, err error)
	DeleteAlamatByID(ctx context.Context, alamatid string) (res string, err error)
}

type AlamatRepositoryImpl struct {
	db *gorm.DB
}

func NewAlamatRepository(db *gorm.DB) AlamatRepository {
	return &AlamatRepositoryImpl{
		db: db,
	}
}
func (a *AlamatRepositoryImpl) GetMyAlamat(ctx context.Context, userid string) (res daos.Alamat, err error) {
	if err := a.db.Debug().WithContext(ctx).First(&res, "id_user", userid).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (a *AlamatRepositoryImpl) GetAlamatByID(ctx context.Context, alamatid string) (res daos.Alamat, err error) {
	if err := a.db.First(&res, alamatid).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (a *AlamatRepositoryImpl) CreateAlamat(ctx context.Context, userid string, data daos.Alamat) (res uint, err error) {
	id, _ := strconv.ParseUint(userid, 0, 0)
	dataAlamat := daos.Alamat{
		Id_user:       uint(id),
		Judul_alamat:  data.Judul_alamat,
		Nama_penerima: data.Nama_penerima,
		No_telp:       data.No_telp,
		Detail_alamat: data.Detail_alamat,
	}

	result := a.db.Create(&dataAlamat).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}

	return dataAlamat.Id, nil
}

func (a *AlamatRepositoryImpl) UpdateAlamatByID(ctx context.Context, alamatid string, data daos.Alamat) (res string, err error) {
	var dataAlamat daos.Alamat
	if err = a.db.Where("id = ? ", alamatid).First(&dataAlamat).WithContext(ctx).Error; err != nil {
		return "Update Alamat failed", gorm.ErrRecordNotFound
	}

	if err := a.db.Model(dataAlamat).Updates(&data).Where("id = ? ", alamatid).Error; err != nil {
		return "Update Alamat failed", err
	}

	return res, nil
}

func (a *AlamatRepositoryImpl) DeleteAlamatByID(ctx context.Context, alamatid string) (res string, err error) {
	var dataAlamat daos.Alamat
	if err = a.db.Where("id = ?", alamatid).First(&dataAlamat).WithContext(ctx).Error; err != nil {
		return "Delete book failed", gorm.ErrRecordNotFound
	}

	if err := a.db.Model(dataAlamat).Delete(&dataAlamat).Error; err != nil {
		return "Delete book failed", err
	}

	return res, nil
}
