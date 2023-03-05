package dto

type (
	CategoryRes struct {
		Id            uint   `json:"id"`
		Nama_category string `json:"nama_category"`
	}

	CategoryReq struct {
		Nama_category string `json:"nama_category" validate:"required"`
	}
)
