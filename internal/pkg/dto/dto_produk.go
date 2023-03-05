package dto

type ProdukReqCreate struct {
	Nama_produk    string `form:"nama_produk" validate:"required"`
	Category_id    uint   `form:"category_id" `
	Toko_id        uint   `form:"toko_id" `
	Harga_reseller string `form:"harga_reseller" validate:"required"`
	Harga_konsumen string `form:"harga_konsumen" validate:"required"`
	Stok           uint   `form:"stok" validate:"required"`
	Deskripsi      string `form:"deskripsi" validate:"required"`
	Photos         string `form:"photos"`
}

type ProdukReqUpdate struct {
	Nama_produk    string `form:"nama_produk,omitempty"`
	Category_id    uint   `form:"category_id,omitempty"`
	Harga_reseller string `form:"harga_reseller,omitempty"`
	Harga_konsumen string `form:"harga_konsumen,omitempty"`
	Stok           uint   `form:"stok,omitempty"`
	Deskripsi      string `form:"deskripsi,omitempty"`
	Photos         string `form:"photos,omitempty"`
}

type ProdukRes struct {
	Nama_produk    string `form:"nama_produk"`
	Category_id    uint   `form:"category_id"`
	Harga_reseller string `form:"harga_reseller"`
	Harga_konsumen string `form:"harga_konsumen"`
	Stok           uint   `form:"stok"`
	Deskripsi      string `form:"deskripsi"`
	Photos         string `form:"photos"`
}

type ProdukFilter struct {
	Nama_produk string `query:"nama_produk"`
	Limit       int    `query:"limit"`
	Page        int    `query:"page"`
}
