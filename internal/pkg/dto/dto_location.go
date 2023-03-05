package dto

type (
	ProvinceList struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}

	CityList struct {
		Id         string `json:"id"`
		Regency_id string `json:"regency_id"`
		Name       string `json:"name"`
	}
)
