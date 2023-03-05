package daos

import "time"

type (
	Toko struct {
		Id         uint `gorm:"autoIncrement"`
		Id_user    uint
		Nama_toko  string
		Url_foto   string
		Updated_at time.Time
		Created_at time.Time
		Produk     Produk       `gorm:"foreignKey:Id_toko"`
		Log_produk []Log_produk `gorm:"foreignKey:Id_toko"`
		Detail_trx []Detail_trx `gorm:"foreignKey:Id_toko"`
	}

	FilterToko struct {
		Name   string
		Limit  int
		Offset int
	}
)
