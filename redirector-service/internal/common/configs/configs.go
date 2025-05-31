package configs

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	SERVER_PORT       string
	SCYLLADB_HOSTS    []string
	SCYLLADB_DATABASE string
	KAFKA_BROKERS     []string
)

const (
	KAFKA_TOPIC_REDIRECTOR = "redirector-service.redirector.events"
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
