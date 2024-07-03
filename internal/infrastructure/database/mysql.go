package database

import (
	"boilerplate-go/internal/domain/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLConnection(dsn string) (*gorm.DB, error) {
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    if err := db.AutoMigrate(&model.User{}, &model.Product{}); err != nil {
		log.Fatal("Failed to auto-migrate: ", err)
	}

    return db, nil
}
