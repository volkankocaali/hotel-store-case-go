package database

import (
	"fmt"
	"github.com/volkankocali/hotel-store-case-go/pkg/config"
	"github.com/volkankocali/hotel-store-case-go/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewMysqlDB(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.MySQLUser, cfg.MySQLPassword, cfg.MySQLHost, cfg.MySQLPort, cfg.MySQLDatabase)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.Set("gorm:table_options", "CHARSET=utf8mb4")
	fmt.Println("Connection Opened to Mysql Database")
	migrate := db.AutoMigrate(
		models.Users{},
		models.Reservation{},
	)

	if err != nil {
		log.Fatalf("Error migrating database: %v", migrate.Error)
	}

	return db, err
}
