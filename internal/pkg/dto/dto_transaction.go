package dto

type (
	TransactionReq struct {
		Method_bayar      string `json:"method_bayar" validate:"required"`
		Alamat_kirim      uint   `json:"alamat_kirim" validate:"required"`
		DetailTransaction []DetailTransactionReq
	}

	TransactionRes struct {
		Id           uint
		Harga_total  uint
		Kode_invoice string
		Method_bayar string
		Alamat_kirim AlamatRes
	}

	DetailTransactionReq struct {
		Product_id uint `json:"product_id"`
		Kuantitas  uint `json:"kuantitas"`
	}

	DetailTransactionRes struct {
		Product     []ProdukRes
		Toko        TokoRes
		Kuantitas   int
		Harga_total int
	}
)
