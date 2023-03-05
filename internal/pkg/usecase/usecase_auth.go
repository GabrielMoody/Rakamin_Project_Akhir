package usecase

import (
	"context"
	"errors"
	"log"
	"time"
	"tugas_akhir_example/internal/daos"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthUseCase interface {
	LoginUser(ctx context.Context, params dto.AuthLogin) (res string, err *helper.ErrorStruct)
	SingUpNewUser(ctx context.Context, params dto.AuthRegister) (res dto.AuthRegisterRes, err *helper.ErrorStruct)
}

type AuthUseCaseImpl struct {
	authRepository repository.AuthRepository
	tokoRepository repository.RepositoryToko
}

func NewAuthUseCase(authRepository repository.AuthRepository, tokoRepository repository.RepositoryToko) AuthUseCase {
	return &AuthUseCaseImpl{
		authRepository: authRepository,
		tokoRepository: tokoRepository,
	}
}

func (a *AuthUseCaseImpl) LoginUser(ctx context.Context, params dto.AuthLogin) (res string, err *helper.ErrorStruct) {
	resRepo, errRepo := a.authRepository.LoginUser(ctx, daos.User{
		Notelp: params.Notelp,
	})

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errors.New("user doesn't exist"),
		}
	}

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(resRepo.Kata_sandi), []byte(params.Kata_sandi)); err != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	signedToken, jwtErr := utils.GenerateJWTToKen(resRepo.Id)

	if jwtErr != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  jwtErr,
		}
	}

	return signedToken, nil
}

func (a *AuthUseCaseImpl) SingUpNewUser(ctx context.Context, params dto.AuthRegister) (res dto.AuthRegisterRes, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(params); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	formatDate := "2001-01-01"
	date, _ := time.Parse(formatDate, params.Tanggal_Lahir)

	hashedPass, errBcrypt := bcrypt.GenerateFromPassword([]byte(params.Kata_sandi), bcrypt.DefaultCost)

	if errBcrypt != nil {
		return res, &helper.ErrorStruct{
			Err:  err.Err,
			Code: fiber.StatusBadRequest,
		}
	}

	resrepo, errRepo := a.authRepository.RegisterNewUser(ctx, daos.User{
		Nama:          params.Nama,
		Notelp:        params.No_telp,
		Email:         params.Email,
		Kata_sandi:    string(hashedPass),
		Tanggal_lahir: &date,
		Jenis_kelamin: params.Jenis_kelamin,
		Pekerjaan:     params.Pekerjaan,
		Id_provinsi:   params.Id_provinsi,
		Id_kota:       params.Id_kota,
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	_, errToko := a.tokoRepository.CreateTokoByUserRegistration(ctx, daos.User{
		Id:   resrepo.Id,
		Nama: resrepo.Nama,
	})

	if errToko != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errToko,
		}
	}

	return dto.AuthRegisterRes{
		Nama:          resrepo.Nama,
		No_telp:       resrepo.Notelp,
		Email:         resrepo.Email,
		Pekerjaan:     resrepo.Pekerjaan,
		Tanggal_Lahir: resrepo.Tanggal_lahir,
		Id_provinsi: dto.Provinsi{
			Id:   params.Id_provinsi,
			Name: "",
		},
		Id_kota: dto.Kota{
			Id:          params.Id_kota,
			Id_provinsi: params.Id_provinsi,
			Name:        "",
		},
	}, nil
}
