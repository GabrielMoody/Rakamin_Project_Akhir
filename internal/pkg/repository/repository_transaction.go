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
	CreateTransaction(ctx context.Context, userid string, produks []uint, params daos.Trx) (res daos.Trx, err error)
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

func (a *TransactionRepositoryImpl) CreateTransaction(ctx context.Context, userid string, produks []uint, params daos.Trx) (res daos.Trx, err error) {
	id, _ := strconv.ParseUint(userid, 0, 0)
	trx := daos.Trx{
		Id_user:           uint(id),
		Method_bayar:      params.Method_bayar,
		Alamat_pengiriman: params.Alamat_pengiriman,
	}

	var detailTrx []daos.Detail_trx
	var produk daos.Produk
	var logProduk daos.Log_produk

	for i, _ := range params.Detail_trx {
		a.db.First(&produk, produks[i])

		logProduk = daos.Log_produk{
			Id_produk:      produks[i],
			Nama_produk:    produk.Nama_produk,
			Slug:           produk.Slug,
			Harga_konsumen: produk.Harga_konsumen,
			Harga_reseller: produk.Harga_reseller,
			Deskripsi:      produk.Deskripsi,
			Id_toko:        produk.Id_toko,
			Id_category:    produk.Id_category,
		}

		a.db.Create(&logProduk)

		harga, _ := strconv.ParseUint(produk.Harga_reseller, 0, 0)
		hargaReseller := params.Detail_trx[i].Harga_total * uint(harga)
		detailTrx = append(detailTrx, daos.Detail_trx{
			Id_trx:        trx.Id,
			Id_log_produk: logProduk.Id,
			Id_toko:       produk.Id_toko,
			Kuantitas:     params.Detail_trx[i].Kuantitas,
			Harga_total:   hargaReseller,
		})

		a.db.Create(&detailTrx)
	}

	if err := a.db.Create(&trx).Error; err != nil {
		return res, err
	}

	// for _, v := range params.Detail_trx {
	// 	harga, _ := strconv.ParseUint(produk.Harga_reseller, 0, 0)
	// 	hargaReseller := v.Kuantitas * uint(harga)
	// 	detailTrx = append(detailTrx, daos.Detail_trx{
	// 		Id_trx:        trx.Id,
	// 		Id_log_produk: logProduk.Id,
	// 		Id_toko:       produk.Id_toko,
	// 		Kuantitas:     v.Kuantitas,
	// 		Harga_total:   hargaReseller,
	// 	})
	// }

	// if err := a.db.Create(&detailTrx).Error; err != nil {
	// 	return res, err
	// }

	return res, nil
}
