package database

import (
	"quotation-server/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnection() (*gorm.DB, error) {
	dsn := "root:root@tcp(localhost:3306)/quotations?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        return nil, err
    }

    db.AutoMigrate(&models.Quotation{})

    return db, nil
}