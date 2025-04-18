package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shortener-service/internal/account/application/services"
	"github.com/shortener-service/internal/account/dal/repositories"
	"github.com/shortener-service/internal/account/interfaces/http"
	"github.com/shortener-service/internal/infrastructure/scylladb"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Logger())

	// -- Initialize the database connection --
	connection := scylladb.NewScyllaDB([]string{"127.0.0.1:9042"}, "shortener")

	// -- Initialize the account handler --
	accountRepository := repositories.NewAccountRepository(connection)
	accountService := services.NewAccountService(accountRepository)
	http.SetupAccountRoutes(r, &http.AccountHandler{AccountService: accountService})

	log.Println("🚀 Server running at http://0.0.0.0:3000")

	for _, route := range r.Routes() {
		log.Println("Registered route:", route.Method, route.Path)
	}

	err := r.Run("0.0.0.0:3000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		return
	}
}
