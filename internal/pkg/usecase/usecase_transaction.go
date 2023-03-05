package usecase

import (
	"context"
	"errors"
	"tugas_akhir_example/internal/daos"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TransactionUseCase interface {
	GetAllTransaction(ctx context.Context) (res []dto.TransactionRes, err *helper.ErrorStruct)
	GetTransactionByID(ctx context.Context, transactionid string) (res dto.TransactionRes, err *helper.ErrorStruct)
	CreateTransaction(ctx context.Context, userid string, params dto.TransactionReq) (res dto.TransactionRes, err *helper.ErrorStruct)
}

type TransactionUseCaseImpl struct {
	transactionrepo repository.TransactionRepository
	alamatrepo      repository.AlamatRepository
}

func NewTransactionUseCase(transactionrepo repository.TransactionRepository, alamatrepo repository.AlamatRepository) TransactionUseCase {
	return &TransactionUseCaseImpl{
		transactionrepo: transactionrepo,
		alamatrepo:      alamatrepo,
	}
}

func (a *TransactionUseCaseImpl) GetAllTransaction(ctx context.Context) (res []dto.TransactionRes, err *helper.ErrorStruct) {
	resRepo, errRepo := a.transactionrepo.GetAllTransaction(ctx)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	for _, v := range resRepo {
		resAlamat, _ := a.alamatrepo.GetAlamatByID(ctx, string(rune(v.Alamat_pengiriman)))

		res = append(res, dto.TransactionRes{
			Id:           v.Id,
			Harga_total:  v.Harga_total,
			Method_bayar: v.Method_bayar,
			Kode_invoice: v.Kode_invoice,
			Alamat_kirim: dto.AlamatRes{
				Id:            resAlamat.Id,
				Judul_alamat:  resAlamat.Judul_alamat,
				Penerima:      resAlamat.Nama_penerima,
				Notelp:        resAlamat.No_telp,
				Detail_alamat: resAlamat.Detail_alamat,
			},
		})
	}

	return res, nil
}

func (a *TransactionUseCaseImpl) GetTransactionByID(ctx context.Context, transactionid string) (res dto.TransactionRes, err *helper.ErrorStruct) {
	resRepo, errRepo := a.transactionrepo.GetTransactionByID(ctx, transactionid)

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Category"),
		}
	}

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	resAlamat, _ := a.alamatrepo.GetAlamatByID(ctx, string(rune(resRepo.Alamat_pengiriman)))

	res = dto.TransactionRes{
		Id:           resRepo.Id,
		Method_bayar: resRepo.Method_bayar,
		Harga_total:  resRepo.Harga_total,
		Kode_invoice: resRepo.Kode_invoice,
		Alamat_kirim: dto.AlamatRes{
			Id:            resAlamat.Id,
			Judul_alamat:  resAlamat.Judul_alamat,
			Penerima:      resAlamat.Nama_penerima,
			Notelp:        resAlamat.No_telp,
			Detail_alamat: resAlamat.Detail_alamat,
		},
	}

	return res, nil
}

func (a *TransactionUseCaseImpl) CreateTransaction(ctx context.Context, userid string, params dto.TransactionReq) (res dto.TransactionRes, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(params); errValidate != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errValidate,
		}
	}

	var detailTrx []daos.Detail_trx
	var produks []uint

	for _, v := range params.DetailTransaction {
		detailTrx = append(detailTrx, daos.Detail_trx{
			Kuantitas: v.Kuantitas,
		})

		produks = append(produks, v.Product_id)
	}

	// for _, v := range params.DetailTransaction {
	// 	produks = append(produks, uint(v.Product_id))
	// }

	resRepo, errRepo := a.transactionrepo.CreateTransaction(ctx, userid, produks, daos.Trx{
		Method_bayar:      params.Method_bayar,
		Alamat_pengiriman: params.Alamat_kirim,
		Detail_trx:        detailTrx,
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return dto.TransactionRes{
		Id:           resRepo.Id,
		Method_bayar: resRepo.Method_bayar,
		Harga_total:  resRepo.Harga_total,
		Kode_invoice: resRepo.Kode_invoice,
	}, nil
}
