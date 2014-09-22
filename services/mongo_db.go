package services

import "labix.org/v2/mgo"

// MongoPool ...
var MongoDB *mgo.Database

// InitMongoDB ...
func InitMongoDB(connectionString string) *mgo.Database {
	var err error

	session, err := mgo.Dial(connectionString)

	if err != nil {
		Logger.Error("Connection to MongoDB failed.")
		panic(err)
	}

	Logger.Info("Connection do MongoDB established sucessfully.")

	MongoDB = session.DB("gonalytics")
	return MongoDB
}
