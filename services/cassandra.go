package services

import (
	"github.com/gocql/gocql"
	"github.com/piotrkowalczuk/gonalytics-backend/lib"
)

// Singleton instance of cassandra session.
var Cassandra *gocql.Session

// InitCassandra ...
func InitCassandra(config lib.CassandraConfig) {
	cluster := gocql.NewCluster(config.GetHosts()...)
	cluster.Keyspace = config.Keyspace
	cluster.Consistency = gocql.Quorum
	cassandra, err := cluster.CreateSession()

	if err != nil {
		Logger.Error("Connection to Cassandra failed.")
		panic(err)
	}

	Logger.Info("Connection do Cassandra established sucessfully.")

	Cassandra = cassandra
}
