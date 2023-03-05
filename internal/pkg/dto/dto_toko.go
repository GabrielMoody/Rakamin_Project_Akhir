package dto

type (
	TokoUpdateReq struct {
		Nama_toko string `json:"nama_toko" validate:"required"`
		Photo     string `json:"photo"`
	}

	TokoRes struct {
		Id        uint   `json:"id"`
		Nama_toko string `json:"nama_toko"`
		Url_photo string `json:"url_foto"`
	}

	FilterToko struct {
		Nama  string `query:"nama"`
		Limit int    `query:"limit"`
		Page  int    `query:"page"`
	}
)
