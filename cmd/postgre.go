package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"inventory-service/internal/domain"
	"log"
)

func NewDB(cfg DbConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	db.AutoMigrate(&domain.Product{}, &domain.Category{})
	return db
}
