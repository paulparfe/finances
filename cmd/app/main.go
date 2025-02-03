package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/paulparfe/finances/internal/adapters/db/postgres"
	"github.com/paulparfe/finances/internal/controller/http/v1"
	"github.com/paulparfe/finances/internal/domain/service"
	transactionusecase "github.com/paulparfe/finances/internal/domain/usecase/transaction"
	userusecase "github.com/paulparfe/finances/internal/domain/usecase/user"
	"log"
	"os"
)

type config struct {
	port string
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg config
	cfg.port = os.Getenv("PORT")

	db := &sql.DB{}

	transactionStorage := postgres.NewTransactionStorage(db)
	transactionService := service.NewTransactionService(transactionStorage)
	transactionUseCase := transactionusecase.NewTransactionUseCase(transactionService)
	transactionHandler := v1.NewTransactionHandler(transactionUseCase)

	userStorage := postgres.NewUserStorage(db)
	userService := service.NewUserService(userStorage)
	userUseCase := userusecase.NewUserUseCase(userService)
	userHandler := v1.NewUserHandler(userUseCase)

	router := gin.Default()
	transactionHandler.Register(router)
	userHandler.Register(router)
	router.Run(cfg.port)
}
