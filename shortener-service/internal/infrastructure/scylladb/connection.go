package scylladb

import (
	"log"
	"time"

	"github.com/gocql/gocql"
)

type ScyllaDB struct {
	Session *gocql.Session
}

func NewScyllaDB(hosts []string, keyspace string) *ScyllaDB {
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Quorum
	cluster.Timeout = 5 * time.Second
	cluster.ConnectTimeout = 5 * time.Second

	// Optional: auth
	// cluster.Authenticator = gocql.PasswordAuthenticator{
	// 	Username: "scylla",
	// 	Password: "secret",
	// }

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to connect to ScyllaDB: %v", err)
	}

	log.Println("Connected to ScyllaDB")

	return &ScyllaDB{Session: session}
}

func (s *ScyllaDB) Close() {
	s.Session.Close()
}
