package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shortener-service/internal/common/configs"
	"github.com/shortener-service/internal/common/logger"
	"github.com/shortener-service/internal/domain/redirector/application/services"
	"github.com/shortener-service/internal/domain/redirector/dal/repositories"
	redirector_http "github.com/shortener-service/internal/domain/redirector/interfaces/http"
	"github.com/shortener-service/internal/infrastructure/kafka"
	"github.com/shortener-service/internal/infrastructure/scylladb"
)

func main() {
	configs.Load()
	logger := logger.NewLogger("redirector-service")

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Logger())

	// -- Initialize the database connection --
	connection := scylladb.NewScyllaDB(configs.SCYLLADB_HOSTS, configs.SCYLLADB_DATABASE)

	// -- Initialize the redirector domain --
	redirectorEventPublisher := kafka.NewKafkaPublisher(
		&kafka.KafkaPublisherConfig{
			Addrs: configs.KAFKA_BROKERS,
			Topic: kafka.Topic{
				Name:       configs.KAFKA_TOPIC_REDIRECTOR,
				Partitions: 1,
			},
		},
	)
	shortURLRepository := repositories.NewShortURLRepository(connection)
	redirectorService := services.NewRedirectorService(shortURLRepository, redirectorEventPublisher)
	redirector_http.SetupRedirectorRoutes(r, &redirector_http.RedirectorHandler{RedirectorService: redirectorService})

	logger.Info("Server running at http://0.0.0.0:" + configs.SERVER_PORT)
	if err := r.Run("0.0.0.0:" + configs.SERVER_PORT); err != nil {
		logger.Error(err.Error())
	}
}
