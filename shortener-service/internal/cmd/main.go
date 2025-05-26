package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shortener-service/internal/common/logger"
	account_service "github.com/shortener-service/internal/domain/account/application/services"
	account_repositories "github.com/shortener-service/internal/domain/account/dal/repositories"
	account_http "github.com/shortener-service/internal/domain/account/interfaces/http"
	shortener_service "github.com/shortener-service/internal/domain/shortener/application/services"
	shortener_repositories "github.com/shortener-service/internal/domain/shortener/dal/repositories"
	shortener_http "github.com/shortener-service/internal/domain/shortener/interfaces/http"
	"github.com/shortener-service/internal/infrastructure/kafka"
	"github.com/shortener-service/internal/infrastructure/scylladb"
)

func main() {
	logger := logger.NewLogger("shortener-service")
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Logger())

	// -- Initialize the database connection --
	connection := scylladb.NewScyllaDB([]string{"127.0.0.1:9042"}, "shortener")

	// -- Initialize the account domain --
	accountEventPublisher := kafka.NewKafkaPublisher(
		&kafka.KafkaPublisherConfig{
			Addrs: []string{"127.0.0.1:29092"},
			Topic: kafka.Topic{
				Name:       "shortener-service.account.events",
				Partitions: 1,
			},
		},
	)
	accountRepository := account_repositories.NewAccountRepository(connection)
	accountService := account_service.NewAccountService(accountRepository, accountEventPublisher)
	account_http.SetupAccountRoutes(r, &account_http.AccountHandler{AccountService: accountService})

	// -- Initialize the shortener domain --
	shortURLRepository := shortener_repositories.NewShortURLRepository(connection)
	shortenerService := shortener_service.NewShortenerService(shortURLRepository)
	shortener_http.SetupShortenerRoutes(r, &shortener_http.ShortenerHandler{ShortenerService: shortenerService})

	logger.Info("Server running at http://0.0.0.0:3000")
	if err := r.Run("0.0.0.0:3000"); err != nil {
		logger.Error(err.Error())
	}
}
