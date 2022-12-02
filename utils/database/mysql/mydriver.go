package mysql

import (
	"fajar/clean/config"
	user "fajar/clean/features/user/repository"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("tidak bisa koneksi ke database anda")
	}

	migrateDB(db)

	return db
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
}
