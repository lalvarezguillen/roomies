package config

import (
	"os"

	"gopkg.in/mgo.v2"
)

type DB struct {
	Session *mgo.Session
}

func (db *DB) DoDial() (s *mgo.Session, err error) {
	return mgo.Dial(DBUrl())
}

func (db *DB) Name() string {
	return "roomies_testing"
}

func DBUrl() string {
	dburl := os.Getenv("MONGO_HOST")

	if dburl == "" {
		dburl = "localhost"
	}

	return dburl
}
