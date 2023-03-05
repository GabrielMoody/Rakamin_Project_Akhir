package dto

type (
	ProfileReqUpdate struct {
		Nama          string `json:"nama" validate:"required"`
		Kata_sandi    string `json:"kata_sandi" validate:"required"`
		No_telp       string `json:"no_telp" validate:"required"`
		Tanggal_Lahir string `json:"tanggal_lahir" validate:"required"`
		Pekerjaan     string `json:"pekerjaan"`
		Email         string `json:"email" validate:"required"`
		Id_provinsi   string `json:"id_provinsi"`
		Id_kota       string `json:"id_kota"`
	}

	ProfileRes struct {
		Nama          string
		No_telp       string
		Tanggal_Lahir string
		Pekerjaan     string
		Email         string
		Id_provinsi   string
		Id_kota       string
	}
)
