package usecase

import (
	"context"
	"errors"
	"log"
	"tugas_akhir_example/internal/daos"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProdukUseCase interface {
	GetAllProduk(ctx context.Context, params dto.ProdukFilter) (res []dto.ProdukRes, err *helper.ErrorStruct)
	GetProdukByID(ctx context.Context, produkid string) (res dto.ProdukRes, err *helper.ErrorStruct)
	CreateProduk(ctx context.Context, data dto.ProdukReqCreate) (res uint, err *helper.ErrorStruct)
	UpdateProdukByID(ctx context.Context, produkid string, data dto.ProdukReqUpdate) (res string, err *helper.ErrorStruct)
	DeleteProdukByID(ctx context.Context, produkid string) (res string, err *helper.ErrorStruct)
}

type ProdukUseCaseImpl struct {
	ProdukRepository repository.ProdukRepository
}

func NewProdukUseCase(produkrepository repository.ProdukRepository) ProdukUseCase {
	return &ProdukUseCaseImpl{
		ProdukRepository: produkrepository,
	}
}

func (a *ProdukUseCaseImpl) GetAllProduk(ctx context.Context, params dto.ProdukFilter) (res []dto.ProdukRes, err *helper.ErrorStruct) {
	if params.Limit < 1 {
		params.Limit = 10
	}

	if params.Page < 1 {
		params.Page = 0
	} else {
		params.Page = (params.Page - 1) * params.Limit
	}

	resRepo, errRepo := a.ProdukRepository.GetAllProduk(ctx, daos.FilterProduk{
		Limit:       params.Limit,
		Offset:      params.Page,
		Nama_produk: params.Nama_produk,
	})

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Produk"),
		}
	}

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	for _, v := range resRepo {
		res = append(res, dto.ProdukRes{
			Nama_produk:    v.Nama_produk,
			Stok:           v.Stock,
			Harga_reseller: v.Harga_reseller,
			Harga_konsumen: v.Harga_konsumen,
			Deskripsi:      v.Deskripsi,
			Category_id:    v.Id_category,
		})
	}

	return res, nil
}

func (a *ProdukUseCaseImpl) GetProdukByID(ctx context.Context, produkid string) (res dto.ProdukRes, err *helper.ErrorStruct) {
	resRepo, errRepo := a.ProdukRepository.GetProdukByID(ctx, produkid)

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Produk"),
		}
	}

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	res = dto.ProdukRes{
		Nama_produk:    resRepo.Nama_produk,
		Harga_reseller: resRepo.Harga_reseller,
		Harga_konsumen: resRepo.Harga_konsumen,
		Stok:           resRepo.Stock,
		Category_id:    resRepo.Id_category,
	}

	return res, nil
}

func (a *ProdukUseCaseImpl) CreateProduk(ctx context.Context, data dto.ProdukReqCreate) (res uint, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := a.ProdukRepository.CreateProduk(ctx, daos.Produk{
		Nama_produk:    data.Nama_produk,
		Harga_reseller: data.Harga_reseller,
		Harga_konsumen: data.Harga_konsumen,
		Id_category:    data.Category_id,
		Id_toko:        data.Toko_id,
		Stock:          data.Stok,
		Deskripsi:      data.Deskripsi,
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}

	return resRepo, nil
}

func (a *ProdukUseCaseImpl) UpdateProdukByID(ctx context.Context, produkid string, data dto.ProdukReqUpdate) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := a.ProdukRepository.UpdateProdukByID(ctx, produkid, daos.Produk{
		Nama_produk:    data.Nama_produk,
		Harga_reseller: data.Harga_reseller,
		Harga_konsumen: data.Harga_konsumen,
		Stock:          data.Stok,
		Deskripsi:      data.Deskripsi,
		Id_category:    data.Category_id,
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}

	return resRepo, nil
}

func (a *ProdukUseCaseImpl) DeleteProdukByID(ctx context.Context, produkid string) (res string, err *helper.ErrorStruct) {
	resRepo, errRepo := a.ProdukRepository.DeleteProdukByID(ctx, produkid)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}

	return resRepo, nil
}
