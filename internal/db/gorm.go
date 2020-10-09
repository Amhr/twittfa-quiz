package db

import (
	"WhoKnowsMeapp/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGorm() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_PORT, config.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	autoMigrate(db)
	return db, nil
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Quiz{})
	db.AutoMigrate(&QuizQuestion{})
	db.AutoMigrate(&Result{})
}
