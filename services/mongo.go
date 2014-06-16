package services

import (
    "labix.org/v2/mgo"
    "github.com/daviddengcn/go-colortext"
     "log"
)

var MongoPool Pool = Pool{}

type Mongo struct {
    connectionString string
    databaseName string
}

func NewMongo(connectionString, databaseName string) (mongo *Mongo) {
    mongo = new(Mongo)
    mongo.connectionString = connectionString
    mongo.databaseName = databaseName

    return
}

func (m *Mongo) Connect() {
	var err error

	if MongoPool.Session, err = mgo.Dial(m.connectionString); err != nil {
		panic(err)
	}

    ct.ChangeColor(ct.Yellow, false, ct.None, false)
	log.Println("Connected to MongoDB.")

	MongoPool.Database = MongoPool.Session.DB(m.databaseName)
}

func (m *Mongo) Disconnect() {
    MongoPool.Session.Close()
}

type Pool struct {
    Session *mgo.Session
    Database *mgo.Database
}

func (p *Pool) Collection(name string) (collection *mgo.Collection) {
    collection = p.Database.C(name)

    return
}
