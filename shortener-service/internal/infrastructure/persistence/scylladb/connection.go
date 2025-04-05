package scylladb

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

type ScyllaDBConfig struct {
	Hosts       []string
	Port        int
	Keyspace    string
	Username    string
	Password    string
	Consistency gocql.Consistency
}

type ScyllaDBClient struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

type ScyllaDBClientInterface interface {
	Session() *gocql.Session
}

func NewScyllaDBClient(cfg ScyllaDBConfig) (*ScyllaDBClient, error) {
	cluster := gocql.NewCluster(cfg.Hosts...)
	cluster.Port = cfg.Port
	cluster.Keyspace = cfg.Keyspace
	cluster.Consistency = cfg.Consistency

	if cfg.Username != "" && cfg.Password != "" {
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: cfg.Username,
			Password: cfg.Password,
		}
	}

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, fmt.Errorf("failed to create ScyllaDB session: %w", err)
	}

	client := &ScyllaDBClient{
		cluster: cluster,
		session: session,
	}

	log.Println("Successfully connected to ScyllaDB cluster.")
	return client, nil
}

func (c *ScyllaDBClient) Session() *gocql.Session {
	return c.session
}
