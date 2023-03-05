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

type TokoUseCase interface {
	GetMyToko(ctx context.Context, userid string) (res dto.TokoRes, err *helper.ErrorStruct)
	GetAllToko(ctx context.Context, params dto.FilterToko) (res []dto.TokoRes, err *helper.ErrorStruct)
	GetTokoByID(ctx context.Context, tokoid string) (res dto.TokoRes, err *helper.ErrorStruct)
	UpdateTokoByID(ctx context.Context, tokoid string, data dto.TokoUpdateReq) (res string, err *helper.ErrorStruct)
}

type TokoUseCaseImpl struct {
	tokoRepository repository.RepositoryToko
}

func NewTokoUseCase(tokoRepo repository.RepositoryToko) TokoUseCase {
	return &TokoUseCaseImpl{
		tokoRepository: tokoRepo,
	}
}

func (a *TokoUseCaseImpl) GetMyToko(ctx context.Context, userid string) (res dto.TokoRes, err *helper.ErrorStruct) {
	resRepo, errRepo := a.tokoRepository.GetMyToko(ctx, userid)

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errors.New("No data toko"),
		}
	}

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	res = dto.TokoRes{
		Id:        resRepo.Id,
		Nama_toko: resRepo.Nama_toko,
		Url_photo: resRepo.Url_foto,
	}

	return res, nil
}

func (a *TokoUseCaseImpl) GetAllToko(ctx context.Context, params dto.FilterToko) (res []dto.TokoRes, err *helper.ErrorStruct) {
	if params.Limit < 1 {
		params.Limit = 10
	}

	if params.Page < 1 {
		params.Page = 0
	} else {
		params.Page = (params.Page - 1) * params.Limit
	}

	resRepo, errRepo := a.tokoRepository.GetAllToko(ctx, daos.FilterToko{
		Limit:  params.Limit,
		Offset: params.Page,
		Name:   params.Nama,
	})

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Toko"),
		}
	}

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	for _, v := range resRepo {
		res = append(res, dto.TokoRes{
			Id:        v.Id,
			Nama_toko: v.Nama_toko,
			Url_photo: v.Url_foto,
		})
	}

	return res, nil
}

func (a *TokoUseCaseImpl) GetTokoByID(ctx context.Context, tokoid string) (res dto.TokoRes, err *helper.ErrorStruct) {
	resRepo, errRepo := a.tokoRepository.GetTokoByID(ctx, tokoid)

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errors.New("No data toko"),
		}
	}

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	res = dto.TokoRes{
		Id:        resRepo.Id,
		Nama_toko: resRepo.Nama_toko,
		Url_photo: resRepo.Url_foto,
	}

	return res, nil
}

func (a *TokoUseCaseImpl) UpdateTokoByID(ctx context.Context, tokoid string, data dto.TokoUpdateReq) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); err != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errValidate,
		}
	}

	resRepo, errRepo := a.tokoRepository.UpdateTokoByID(ctx, tokoid, daos.Toko{
		Nama_toko: data.Nama_toko,
		Url_foto:  data.Photo,
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
