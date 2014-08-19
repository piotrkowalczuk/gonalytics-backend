package services

import (
	"github.com/fatih/color"
	"labix.org/v2/mgo"
)

// MongoPool ...
var MongoPool = Pool{}

// Mongo ...
type Mongo struct {
	connectionString string
	databaseName     string
}

// NewMongo ...
func NewMongo(connectionString, databaseName string) (mongo *Mongo) {
	mongo = new(Mongo)
	mongo.connectionString = connectionString
	mongo.databaseName = databaseName

	return
}

// Connect ...
func (m *Mongo) Connect() {
	var err error

	if MongoPool.Session, err = mgo.Dial(m.connectionString); err != nil {
		panic(err)
	}

	color.Green("Connected to MongoDB.")
	MongoPool.Database = MongoPool.Session.DB(m.databaseName)
	MongoPool.AddIndexes()
}

// AddIndexes @TODO fix that and put this code to the repositores
func (mp *Pool) AddIndexes() {
	index1 := mgo.Index{
		Key:        []string{"first_action_at_bucket"},
		Unique:     false,
		Background: true,
	}

	index2 := mgo.Index{
		Key:        []string{"last_action_at_bucket"},
		Unique:     false,
		Background: true,
	}

	err := mp.Collection("visit").EnsureIndex(index1)
	err = mp.Collection("visit").EnsureIndex(index2)
	if err != nil {
		panic(err)
	}
}

// Disconnect ...
func (m *Mongo) Disconnect() {
	MongoPool.Session.Close()
}

// Pool ...
type Pool struct {
	Session  *mgo.Session
	Database *mgo.Database
}

// Collection ...
func (p *Pool) Collection(name string) (collection *mgo.Collection) {
	collection = p.Database.C(name)

	return
}
