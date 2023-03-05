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

type AlamatUseCase interface {
	GetMyAlamat(ctx context.Context, userid string) (res dto.AlamatRes, err *helper.ErrorStruct)
	GetAlamatByID(ctx context.Context, alamatid string) (res dto.AlamatRes, err *helper.ErrorStruct)
	CreateAlamat(ctx context.Context, userid string, data dto.AlamatReq) (res uint, err *helper.ErrorStruct)
	UpdateAlamatByID(ctx context.Context, alamatid string, data dto.AlamatReq) (res string, err *helper.ErrorStruct)
	DeleteAlamatByID(ctx context.Context, alamatid string) (res string, err *helper.ErrorStruct)
}

type AlamatUseCaseImpl struct {
	AlamatRepository repository.AlamatRepository
}

func NewAlamatUseCase(Alamatrepository repository.AlamatRepository) AlamatUseCase {
	return &AlamatUseCaseImpl{
		AlamatRepository: Alamatrepository,
	}
}

func (a *AlamatUseCaseImpl) GetMyAlamat(ctx context.Context, userid string) (res dto.AlamatRes, err *helper.ErrorStruct) {
	resRepo, errRepo := a.AlamatRepository.GetMyAlamat(ctx, userid)

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Alamat"),
		}
	}

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return dto.AlamatRes{
		Id:            resRepo.Id,
		Judul_alamat:  resRepo.Judul_alamat,
		Penerima:      resRepo.Nama_penerima,
		Notelp:        resRepo.No_telp,
		Detail_alamat: resRepo.Detail_alamat,
	}, nil
}

func (a *AlamatUseCaseImpl) GetAlamatByID(ctx context.Context, alamatid string) (res dto.AlamatRes, err *helper.ErrorStruct) {
	resRepo, errRepo := a.AlamatRepository.GetAlamatByID(ctx, string(alamatid))

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Alamat"),
		}
	}

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return dto.AlamatRes{
		Id:            resRepo.Id,
		Judul_alamat:  resRepo.Judul_alamat,
		Penerima:      resRepo.Nama_penerima,
		Notelp:        resRepo.No_telp,
		Detail_alamat: resRepo.Detail_alamat,
	}, nil
}

func (a *AlamatUseCaseImpl) CreateAlamat(ctx context.Context, userid string, data dto.AlamatReq) (res uint, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := a.AlamatRepository.CreateAlamat(ctx, userid, daos.Alamat{
		Judul_alamat:  data.Judul_alamat,
		Nama_penerima: data.Penerima,
		No_telp:       data.Notelp,
		Detail_alamat: data.Detail_alamat,
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}

	return resRepo, nil
}

func (a *AlamatUseCaseImpl) UpdateAlamatByID(ctx context.Context, alamatid string, data dto.AlamatReq) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := a.AlamatRepository.UpdateAlamatByID(ctx, alamatid, daos.Alamat{
		Judul_alamat:  data.Judul_alamat,
		Nama_penerima: data.Penerima,
		No_telp:       data.Notelp,
		Detail_alamat: data.Detail_alamat,
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}

	return resRepo, nil
}

func (a *AlamatUseCaseImpl) DeleteAlamatByID(ctx context.Context, alamatid string) (res string, err *helper.ErrorStruct) {
	resRepo, errRepo := a.AlamatRepository.DeleteAlamatByID(ctx, alamatid)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}

	return resRepo, nil
}
