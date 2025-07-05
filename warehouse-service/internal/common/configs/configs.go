package configs

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	SERVER_PORT          string
	CLICK_HOUSE_HOSTS    []string
	CLICK_HOUSE_USER     string
	CLICK_HOUSE_PASS     string
	CLICK_HOUSE_DATABASE string
	KAFKA_BROKERS        []string
)

const (
	KAFKA_GROUP_ACCOUNT   = "warehouse-service.account.events"
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

	CLICK_HOUSE_HOSTS = strings.Split(os.Getenv("CLICK_HOUSE_HOSTS"), ",")
	CLICK_HOUSE_USER = os.Getenv("CLICK_HOUSE_USER")
	CLICK_HOUSE_PASS = os.Getenv("CLICK_HOUSE_PASS")
	CLICK_HOUSE_DATABASE = os.Getenv("CLICK_HOUSE_DATABASE")
}
