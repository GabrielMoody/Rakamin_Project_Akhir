package daos

import "time"

type Category struct {
	Id            uint `gorm:"autoIncrement"`
	Nama_category string
	Updated_at    time.Time
	Created_at    time.Time
	Produk        Produk     `gorm:"foreignKey:Id_category"`
	Log_produk    Log_produk `gorm:"foreignKey:Id_category"`
}
