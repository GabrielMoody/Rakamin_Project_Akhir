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

type CategoryUseCase interface {
	GetAllCategory(ctx context.Context) (res []dto.CategoryRes, err *helper.ErrorStruct)
	GetCategoryByID(ctx context.Context, categoryid string) (res dto.CategoryRes, err *helper.ErrorStruct)
	CreateCategory(ctx context.Context, data dto.CategoryReq) (res uint, err *helper.ErrorStruct)
	UpdateCategoryByID(ctx context.Context, categoryid string, data dto.CategoryReq) (res string, err *helper.ErrorStruct)
	DeleteCategoryByID(ctx context.Context, categoryid string) (res string, err *helper.ErrorStruct)
}

type CategoryUseCaseImpl struct {
	CategoryRepository repository.CategoryRepository
}

func NewCategoryUseCase(Categoryrepository repository.CategoryRepository) CategoryUseCase {
	return &CategoryUseCaseImpl{
		CategoryRepository: Categoryrepository,
	}
}

func (a *CategoryUseCaseImpl) GetAllCategory(ctx context.Context) (res []dto.CategoryRes, err *helper.ErrorStruct) {
	// if params.Limit < 1 {
	// 	params.Limit = 10
	// }

	// if params.Page < 1 {
	// 	params.Page = 0
	// } else {
	// 	params.Page = (params.Page - 1) * params.Limit
	// }

	resRepo, errRepo := a.CategoryRepository.GetAllCategory(ctx)

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

	for _, v := range resRepo {
		res = append(res, dto.CategoryRes{
			Id:            v.Id,
			Nama_category: v.Nama_category,
		})
	}

	return res, nil
}

func (a *CategoryUseCaseImpl) GetCategoryByID(ctx context.Context, categoryid string) (res dto.CategoryRes, err *helper.ErrorStruct) {
	resRepo, errRepo := a.CategoryRepository.GetCategoryByID(ctx, categoryid)

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

	res = dto.CategoryRes{
		Id:            resRepo.Id,
		Nama_category: resRepo.Nama_category,
	}

	return res, nil
}

func (a *CategoryUseCaseImpl) CreateCategory(ctx context.Context, data dto.CategoryReq) (res uint, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := a.CategoryRepository.CreateCategory(ctx, daos.Category{
		Nama_category: data.Nama_category,
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}

	return resRepo, nil
}

func (a *CategoryUseCaseImpl) UpdateCategoryByID(ctx context.Context, categoryid string, data dto.CategoryReq) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := a.CategoryRepository.UpdateCategoryByID(ctx, categoryid, daos.Category{
		Nama_category: data.Nama_category,
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}

	return resRepo, nil
}

func (a *CategoryUseCaseImpl) DeleteCategoryByID(ctx context.Context, categoryid string) (res string, err *helper.ErrorStruct) {
	resRepo, errRepo := a.CategoryRepository.DeleteCategoryByID(ctx, categoryid)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}

	return resRepo, nil
}
