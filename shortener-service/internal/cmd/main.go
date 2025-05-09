package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	account_service "github.com/shortener-service/internal/domain/account/application/services"
	account_repositories "github.com/shortener-service/internal/domain/account/dal/repositories"
	account_http "github.com/shortener-service/internal/domain/account/interfaces/http"
	shortener_service "github.com/shortener-service/internal/domain/shortener/application/services"
	shortener_repositories "github.com/shortener-service/internal/domain/shortener/dal/repositories"
	shortener_http "github.com/shortener-service/internal/domain/shortener/interfaces/http"
	"github.com/shortener-service/internal/infrastructure/scylladb"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Logger())

	// -- Initialize the database connection --
	connection := scylladb.NewScyllaDB([]string{"127.0.0.1:9042"}, "shortener")

	// -- Initialize the account handler --
	accountRepository := account_repositories.NewAccountRepository(connection)
	accountService := account_service.NewAccountService(accountRepository)
	account_http.SetupAccountRoutes(r, &account_http.AccountHandler{AccountService: accountService})

	shortURLRepository := shortener_repositories.NewShortURLRepository(connection)
	shortenerService := shortener_service.NewShortenerService(shortURLRepository)
	shortener_http.SetupShortenerRoutes(r, &shortener_http.ShortenerHandler{ShortenerService: shortenerService})

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
