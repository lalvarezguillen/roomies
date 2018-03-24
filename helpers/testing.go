package helpers

import (
	"github.com/lalvarezguillen/roomies/config"
	"gopkg.in/mgo.v2/bson"
)

// Deletes all the entries in a DB's collection given the name of the collection
func ClearCollection(collName string) {
	db := config.DB{}
	sess, err := db.DoDial()
	if err != nil {
		panic("There was a problem connecting to the DB")
	}
	defer sess.Close()
	coll := sess.DB(db.Name()).C(collName)
	coll.RemoveAll(bson.M{})
}
