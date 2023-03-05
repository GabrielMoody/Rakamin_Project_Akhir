package repository

import (
	"context"
	"strconv"
	"tugas_akhir_example/internal/daos"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetAllTransaction(ctx context.Context) (res []daos.Trx, err error)
	GetTransactionByID(ctx context.Context, transactionid string) (res daos.Trx, err error)
	CreateTransaction(ctx context.Context, userid string, params daos.Trx) (res daos.Trx, err error)
}

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImpl{
		db: db,
	}
}

func (a *TransactionRepositoryImpl) GetAllTransaction(ctx context.Context) (res []daos.Trx, err error) {
	if err := a.db.Find(&res).WithContext(ctx).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (a *TransactionRepositoryImpl) GetTransactionByID(ctx context.Context, transactionid string) (res daos.Trx, err error) {
	if err := a.db.Debug().WithContext(ctx).First(&res, transactionid).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (a *TransactionRepositoryImpl) CreateTransaction(ctx context.Context, userid string, params daos.Trx) (res daos.Trx, err error) {
	id, _ := strconv.ParseUint(userid, 0, 0)
	trx := daos.Trx{
		Id_user:           uint(id),
		Method_bayar:      params.Method_bayar,
		Alamat_pengiriman: params.Alamat_pengiriman,
		Detail_trx:        params.Detail_trx,
	}

	var detailTrx []daos.Detail_trx

	for _, v := range params.Detail_trx {
		detailTrx = append(detailTrx, daos.Detail_trx{
			Id_log_produk: v.Id_log_produk,
			Kuantitas:     v.Kuantitas,
		})
	}

	if err := a.db.Create(&trx).Error; err != nil {
		return res, err
	}

	if err := a.db.Create(&detailTrx).Error; err != nil {
		return res, err
	}

	return res, nil
}
