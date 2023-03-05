package dto

type (
	ProvinceList struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}

	CityList struct {
		Id          string `json:"id"`
		Province_id string `json:"province_id"`
		Name        string `json:"name"`
	}
)
