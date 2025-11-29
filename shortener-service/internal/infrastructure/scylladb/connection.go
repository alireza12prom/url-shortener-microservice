package scylladb

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/shortener-service/internal/common/logger"
)

type ScyllaDB struct {
	session *gocql.Session
	logger  *logger.Logger
}

func NewScyllaDB(hosts []string, keyspace string) *ScyllaDB {
	cluster := gocql.NewCluster(hosts...)
	logger := logger.NewLogger("ScyllaDB")

	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Quorum
	cluster.Timeout = 5 * time.Second
	cluster.ConnectTimeout = 5 * time.Second

	session, err := cluster.CreateSession()
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Info("Connected")

	return &ScyllaDB{session: session, logger: logger}
}

func (db *ScyllaDB) GetSession() *gocql.Session {
	return db.session
}

func (db *ScyllaDB) Close() {
	db.session.Close()
	db.logger.Info("Closed")
}
