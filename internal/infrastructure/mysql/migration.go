package mysql

import (
	"fmt"
	"tugas_akhir_example/internal/daos"
	"tugas_akhir_example/internal/helper"

	"gorm.io/gorm"
)

func RunMigration(mysqlDB *gorm.DB) {

	err := mysqlDB.AutoMigrate(
		&daos.User{},
		&daos.Toko{},
		&daos.Alamat{},
		&daos.Category{},
		&daos.Produk{},
		&daos.Log_produk{},
		&daos.Foto_produk{},
		&daos.Trx{},
		&daos.Detail_trx{},
	)

	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Failed Database Migrated : %s", err.Error()))
	}

	helper.Logger(currentfilepath, helper.LoggerLevelInfo, "Database Migrated")
}
