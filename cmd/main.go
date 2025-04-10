package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"inventory-service/internal/api/http/v1"
	"inventory-service/internal/repo/postgres"
	"inventory-service/usecase"
	"log"
)

func main() {
	err := godotenv.Load(".env")
    	cfg, err := ReadEnv()
    	if err != nil {
    		log.Fatal(err)
    	}
	db := NewDB(cfg.DbConfig)

	productRepo := postgres.NewProductRepo(db)
	productUC := usecase.NewProductUsecase(productRepo)

	categoryRepo := postgres.NewCategoryRepo(db)
	categoryUC := usecase.NewCategoryUsecase(categoryRepo)

	r := v1.NewRouter(productUC, categoryUC)
	err = r.Run(fmt.Sprintf(":%s",cfg.HTTPPort))
	if err != nil {
		log.Fatal(err)
	}
}
