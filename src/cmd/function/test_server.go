package main

import (
	"fmt"
	"net/http"

	"github.com/tora0091/get-account-data/domain"
	"github.com/tora0091/get-account-data/domain/repositories"
	"github.com/tora0091/get-account-data/infra/handler"
	"github.com/tora0091/get-account-data/infra/router"
)

func main() {
	db := domain.NewSqlHandler()

	accountRepo := repositories.NewAccountRepository(db)
	accountHandler := handler.NewAccountHandler(accountRepo)

	route := router.NewRouter(accountHandler)

	http.HandleFunc("/", route.SetRouting)

	fmt.Println("Start Server...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
