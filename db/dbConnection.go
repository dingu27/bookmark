package db

import (
	"os"

	"gopkg.in/mgo.v2"
)

var (
	database = "rr_bookmark"
	host     = os.Getenv("MONGO_HOST")
)

// GetMongoCon database connection method
func GetMongoCon() (*mgo.Database, error) {
	session, err := mgo.Dial(host)

	if err != nil {
		return nil, err
	}
	db := session.DB(database)

	return db, nil
}
