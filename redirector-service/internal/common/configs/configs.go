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
)

func Load() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err.Error())
	}

	SERVER_PORT = os.Getenv("SERVER_PORT")

	SCYLLADB_HOSTS = strings.Split(os.Getenv("SCYLLADB_HOSTS"), ",")
	SCYLLADB_DATABASE = os.Getenv("SCYLLADB_DATABASE")
}
