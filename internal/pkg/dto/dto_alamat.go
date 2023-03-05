package dto

type (
	AlamatReq struct {
		Judul_alamat  string `json:"judul_alamat"`
		Penerima      string `json:"nama_penerima"`
		Notelp        string `json:"no_telp"`
		Detail_alamat string `json:"detail_alamat"`
	}

	AlamatRes struct {
		Id            uint
		Judul_alamat  string
		Penerima      string
		Notelp        string
		Detail_alamat string
	}
)
