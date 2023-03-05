package daos

import "time"

type (
	Produk struct {
		Id             uint `gorm:"autoIncrement"`
		Nama_produk    string
		Slug           string
		Harga_reseller string
		Harga_konsumen string
		Stock          uint
		Deskripsi      string `gorm:"type:text"`
		Updated_at     time.Time
		Created_at     time.Time
		Id_toko        uint
		Id_category    uint
		Foto_produk    []Foto_produk `gorm:"foreignKey:Id_produk"`
		Log_produk     Log_produk    `gorm:"foreignKey:Id_produk"`
	}

	Log_produk struct {
		Id             uint `gorm:"autoIncrement"`
		Id_produk      uint
		Nama_produk    string
		Slug           string
		Harga_reseller string
		Harga_konsumen string
		Deskripsi      string `gorm:"type:text"`
		Created_at     time.Time
		Updated_at     time.Time
		Id_toko        uint
		Id_category    uint
		Detail_trx     Detail_trx `gorm:"foreignKey:id_log_produk"`
	}

	Foto_produk struct {
		Id         uint `gorm:"autoIncrement"`
		Id_produk  uint
		Url        string
		Updated_at time.Time
		Created_at time.Time
	}

	FilterProduk struct {
		Limit       int
		Offset      int
		Nama_produk string
	}
)
