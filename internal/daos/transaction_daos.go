package daos

import "time"

type Trx struct {
	Id                uint `gorm:"autoIncrement"`
	Id_user           uint
	Alamat_pengiriman uint
	Harga_total       uint
	Kode_invoice      string
	Method_bayar      string
	Updated_at        time.Time
	Created_at        time.Time
	Detail_trx        []Detail_trx `gorm:"foreignKey:Id_trx"`
}

type Detail_trx struct {
	Id            uint `gorm:"autoIncrement"`
	Id_trx        uint
	Id_log_produk uint
	Id_toko       uint
	Kuantitas     uint
	Harga_total   uint
	Updated_at    time.Time
	Created_at    time.Time
}

// type TrxFilter struct {

// }
