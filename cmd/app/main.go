package main

import (
	"database/sql"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/paulparfe/finances/internal/adapters/db/postgres"
	"github.com/paulparfe/finances/internal/controller/http/v1"
	"github.com/paulparfe/finances/internal/domain/service"
	transactionusecase "github.com/paulparfe/finances/internal/domain/usecase/transaction"
	userusecase "github.com/paulparfe/finances/internal/domain/usecase/user"
	"os"
)

type config struct {
	port string
}

func main() {

	godotenv.Load(".env", ".env.example")

	var cfg config
	flag.StringVar(&cfg.port, "PORT", os.Getenv("PORT"), "API server port")

	flag.Parse()

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
