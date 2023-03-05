package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"

	"github.com/gofiber/fiber/v2"
)

type ProvinceCityUseCase interface {
	GetListProvincies(ctx context.Context) (res []dto.ProvinceList, err *helper.ErrorStruct)
	GetListCities(ctx context.Context, provinceid string) (res []dto.CityList, err *helper.ErrorStruct)
	GetDetailProvinceByID(ctx context.Context, provinceid string) (res dto.ProvinceList, err *helper.ErrorStruct)
	GetDetailCityByID(ctx context.Context, cityid string) (res dto.CityList, err *helper.ErrorStruct)
}

type ProvinceCityUseCaseImpl struct {
}

func NewProvinceCityUseCase() ProvinceCityUseCase {
	return &ProdukUseCaseImpl{}
}

func (a *ProdukUseCaseImpl) GetListProvincies(ctx context.Context) (res []dto.ProvinceList, err *helper.ErrorStruct) {
	resp, errResp := http.Get("https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json")

	if errResp != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errResp,
		}
	}

	body, errBody := ioutil.ReadAll(resp.Body)

	if errBody != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errBody,
		}
	}

	if errJSON := json.Unmarshal(body, &res); errJSON != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errJSON,
		}
	}

	return res, nil
}

func (a *ProdukUseCaseImpl) GetListCities(ctx context.Context, provinceid string) (res []dto.CityList, err *helper.ErrorStruct) {
	req := fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/regencies/%s.json", provinceid)
	log.Println(req)
	resp, errResp := http.Get(req)

	if errResp != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errResp,
		}
	}

	body, errBody := ioutil.ReadAll(resp.Body)

	if errBody != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errBody,
		}
	}

	if errJSON := json.Unmarshal(body, &res); errJSON != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errJSON,
		}
	}

	return res, nil
}

func (a *ProdukUseCaseImpl) GetDetailProvinceByID(ctx context.Context, provinceid string) (res dto.ProvinceList, err *helper.ErrorStruct) {
	resp, errResp := http.Get("https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json")

	if errResp != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errResp,
		}
	}

	body, errBody := ioutil.ReadAll(resp.Body)

	if errBody != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errBody,
		}
	}

	var list []dto.ProvinceList

	if errJSON := json.Unmarshal(body, &list); errJSON != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errJSON,
		}
	}

	for _, v := range list {
		if v.Id == provinceid {
			return dto.ProvinceList{
				Id:   v.Id,
				Name: v.Name,
			}, nil
		}
	}

	return res, nil
}

func (a *ProdukUseCaseImpl) GetDetailCityByID(ctx context.Context, cityid string) (res dto.CityList, err *helper.ErrorStruct) {
	req := fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/regencies/%s.json", cityid[:2])
	resp, errResp := http.Get(req)

	if errResp != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errResp,
		}
	}

	body, errBody := ioutil.ReadAll(resp.Body)

	if errBody != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errBody,
		}
	}

	var list []dto.CityList

	if errJSON := json.Unmarshal(body, &list); errJSON != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errJSON,
		}
	}

	for _, v := range list {
		if v.Id == cityid {
			return dto.CityList{
				Id:          v.Id,
				Province_id: v.Province_id,
				Name:        v.Name,
			}, nil
		}
	}

	return res, nil
}
