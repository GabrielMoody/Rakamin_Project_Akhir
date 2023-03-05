package usecase

import (
	"context"
	"log"
	"time"
	"tugas_akhir_example/internal/daos"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/repository"

	"github.com/gofiber/fiber/v2"
)

type UserUseCase interface {
	GetMyProfie(ctx context.Context, userid string) (res dto.ProfileRes, err *helper.ErrorStruct)
	UpdateMyProfile(ctx context.Context, userid string, data dto.ProfileReqUpdate) (res dto.ProfileRes, err *helper.ErrorStruct)
}

type UserUseCaseImpl struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &UserUseCaseImpl{
		userRepo: userRepo,
	}
}

func (a *UserUseCaseImpl) GetMyProfie(ctx context.Context, userid string) (res dto.ProfileRes, err *helper.ErrorStruct) {
	resRepo, errRepo := a.userRepo.GetMyProfile(ctx, userid)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return dto.ProfileRes{
		Nama:          resRepo.Nama,
		No_telp:       resRepo.Notelp,
		Tanggal_Lahir: resRepo.Tanggal_lahir.String(),
		Pekerjaan:     resRepo.Pekerjaan,
		Email:         resRepo.Email,
		Id_provinsi:   resRepo.Id_provinsi,
		Id_kota:       resRepo.Id_kota,
	}, nil
}

func (a *UserUseCaseImpl) UpdateMyProfile(ctx context.Context, userid string, data dto.ProfileReqUpdate) (res dto.ProfileRes, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	date, _ := time.Parse("2020-01-01", data.Tanggal_Lahir)

	resRepo, errRepo := a.userRepo.UpdateMyProfile(ctx, userid, daos.User{
		Nama:          data.Nama,
		Notelp:        data.No_telp,
		Pekerjaan:     data.Pekerjaan,
		Email:         data.Email,
		Kata_sandi:    data.Kata_sandi,
		Tanggal_lahir: &date,
		Id_provinsi:   data.Id_provinsi,
		Id_kota:       data.Id_kota,
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}

	return dto.ProfileRes{
		Nama:          resRepo.Nama,
		No_telp:       resRepo.Notelp,
		Tanggal_Lahir: resRepo.Tanggal_lahir.String(),
		Pekerjaan:     resRepo.Pekerjaan,
		Email:         resRepo.Email,
		Id_provinsi:   resRepo.Id_provinsi,
		Id_kota:       resRepo.Id_kota,
	}, nil
}
