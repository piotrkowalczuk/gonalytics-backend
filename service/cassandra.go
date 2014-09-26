package service

import (
	"github.com/gocql/gocql"
)

// MongoPool ...
var Cassandra *gocql.Session

// InitCassandra ...
func InitCassandra(keyspace string, addresses []string) *gocql.Session {
	cluster := gocql.NewCluster(addresses...)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Quorum
	Cassandra, err := cluster.CreateSession()

	if err != nil {
		Logger.Error("Connection to Cassandra failed.")
		panic(err)
	}

	Logger.Info("Connection do Cassandra established sucessfully.")

	return Cassandra
}
