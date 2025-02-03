package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/paulparfe/finances/internal/adapters/db/postgres"
	"github.com/paulparfe/finances/internal/controller/http/v1"
	"github.com/paulparfe/finances/internal/domain/service"
	transactionusecase "github.com/paulparfe/finances/internal/domain/usecase/transaction"
	userusecase "github.com/paulparfe/finances/internal/domain/usecase/user"
	postgresql "github.com/paulparfe/finances/pkg/client/postgresql"
	"log"
	"os"
	"time"
)

type config struct {
	port string
	DB   postgresql.TmpDBConfig
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg config
	cfg.port = os.Getenv("PORT")
	cfg.DB.Host = os.Getenv("DB_HOST")
	cfg.DB.Port = os.Getenv("DB_PORT")
	cfg.DB.Database = os.Getenv("DB_DATABASE")
	cfg.DB.Username = os.Getenv("DB_USERNAME")
	cfg.DB.Password = os.Getenv("DB_PASSWORD")

	ctx := context.Background()
	postgreSQLClient, err := postgresql.NewClient(ctx, 30, 2*time.Second, cfg.DB)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer postgreSQLClient.Close()

	log.Println("Connected to the database successfully!")

	transactionStorage := postgres.NewTransactionStorage(postgreSQLClient)
	transactionService := service.NewTransactionService(transactionStorage)
	transactionUseCase := transactionusecase.NewTransactionUseCase(transactionService)
	transactionHandler := v1.NewTransactionHandler(transactionUseCase)

	userStorage := postgres.NewUserStorage(postgreSQLClient)
	userService := service.NewUserService(userStorage)
	userUseCase := userusecase.NewUserUseCase(userService)
	userHandler := v1.NewUserHandler(userUseCase)

	router := gin.Default()
	transactionHandler.Register(router)
	userHandler.Register(router)
	router.Run(cfg.port)
}
