package database

import (
	"quotation-server/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func OpenConnection() (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open("quotation.db"), &gorm.Config{})

    if err != nil {
        return nil, err
    }

    db.AutoMigrate(&models.Quotation{})

    return db, nil
}