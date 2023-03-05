package dto

import "time"

type (
	AuthLogin struct {
		Notelp     string `json:"no_telp" validate:"required"`
		Kata_sandi string `json:"kata_sandi" validate:"required"`
	}

	AuthRegister struct {
		Nama          string `json:"nama" validate:"required"`
		Kata_sandi    string `json:"kata_sandi" validate:"required"`
		No_telp       string `json:"no_telp" validate:"required"`
		Tanggal_Lahir string `json:"tanggal_lahir" validate:"required"`
		Jenis_kelamin string `json:"jenis_kelamin"`
		Pekerjaan     string `json:"pekerjaan"`
		Email         string `json:"email" validate:"required"`
		Id_provinsi   string `json:"id_provinsi"`
		Id_kota       string `json:"id_kota"`
	}

	AuthRegisterRes struct {
		Nama          string     `json:"nama"`
		Kata_sandi    string     `json:"kata_sandi"`
		No_telp       string     `json:"no_telp"`
		Tanggal_Lahir *time.Time `json:"tanggal_lahir"`
		Pekerjaan     string     `json:"pekerjaan"`
		Email         string     `json:"email"`
		Id_provinsi   Provinsi   `json:"id_provinsi"`
		Id_kota       Kota       `json:"id_kota"`
	}

	Provinsi struct {
		Id   string
		Name string
	}

	Kota struct {
		Id          string
		Id_provinsi string
		Name        string
	}
)
