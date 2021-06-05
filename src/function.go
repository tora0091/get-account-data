package main

import (
	"github.com/tora0091/get-account-data/domain"
	"github.com/tora0091/get-account-data/domain/repositories"
	"github.com/tora0091/get-account-data/infra/handler"
	"github.com/tora0091/get-account-data/infra/router"
)

func GetAccountData() {
	db := domain.NewSqlHandler()

	accountRepo := repositories.NewAccountRepository(db)
	accountHandler := handler.NewAccountHandler(accountRepo)

	router.NewRouter(accountHandler)
}
