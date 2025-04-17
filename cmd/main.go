package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"inventory-service/internal/api/grpc/handlers"
	"inventory-service/internal/repo/postgres"
	"inventory-service/internal/usecase"
	"inventory-service/proto/inventorypb"
	"log"
	"net"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg, err := ReadEnv()
	if err != nil {
		log.Fatal(err)
	}
	db := NewDB(cfg.DbConfig)

	productRepo := postgres.NewProductRepo(db)
	productUC := usecase.NewProductUsecase(productRepo)

	categoryRepo := postgres.NewCategoryRepo(db)
	categoryUC := usecase.NewCategoryUsecase(categoryRepo)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPCPort)) // в .env добавь GRPCPort=50051
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Регистрируем хэндлер, реализующий InventoryServiceServer
	inventorypb.RegisterInventoryServiceServer(grpcServer, handlers.NewInventoryHandler(categoryUC, productUC))

	log.Printf(" gRPC-сервер запущен на порту %s\n", cfg.GRPCPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}

	/*r := v1.NewRouter(productUC, categoryUC)
	err = r.Run(fmt.Sprintf(":%s", cfg.HTTPPort))
	if err != nil {
		log.Fatal(err)
	}*/

}
