package daos

import "time"

type User struct {
	Id            uint `gorm:"autoIncrement"`
	Nama          string
	Kata_sandi    string
	Notelp        string     `gorm:"unique"`
	Tanggal_lahir *time.Time `gorm:"type:date"`
	Jenis_kelamin string
	Tentang       string `gorm:"type:text"`
	Pekerjaan     string
	Email         string `gorm:"unique"`
	Id_provinsi   string
	Id_kota       string
	IsAdmin       bool
	Updated_at    time.Time
	Created_at    time.Time
	Alamat        []Alamat `gorm:"foreignKey:Id_user"`
	Trx           Trx      `gorm:"foreignKey:Id_user"`
	Toko          Toko     `gorm:"foreignKey:Id_user"`
}
