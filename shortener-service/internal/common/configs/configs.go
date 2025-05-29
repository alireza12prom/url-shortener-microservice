package configs

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	SERVER_PORT       string
	KAFKA_BROKERS     []string
	SCYLLADB_HOSTS    []string
	SCYLLADB_DATABASE string
)

const (
	KAFKA_TOPIC_ACCOUNT   = "shortener-service.account.events"
	KAFKA_TOPIC_SHORTENER = "shortener-service.shortener.events"
)

func Load() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err.Error())
	}

	SERVER_PORT = os.Getenv("SERVER_PORT")

	KAFKA_BROKERS = strings.Split(os.Getenv("KAFKA_BROKERS"), ",")

	SCYLLADB_HOSTS = strings.Split(os.Getenv("SCYLLADB_HOSTS"), ",")
	SCYLLADB_DATABASE = os.Getenv("SCYLLADB_DATABASE")
}
