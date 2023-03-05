package daos

import "time"

type Alamat struct {
	Id            uint `gorm:"autoIncrement"`
	Id_user       uint
	Judul_alamat  string
	Nama_penerima string
	No_telp       string
	Detail_alamat string
	Updated_at    time.Time
	Created_at    time.Time
	Trx           Trx `gorm:"foreignKey:Alamat_pengiriman"`
}
