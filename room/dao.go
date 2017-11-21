package room

import (
	"errors"
	"time"

	"github.com/lalvarezguillen/roomies/config"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

// GetByID fetches a Room from DB by ID
func GetByID(id string) (*Room, error) {
	db := config.DB{}
	sess, err := db.DoDial()
	r := &Room{}
	if err != nil {
		return r, errors.New("There was a problem connecting to the DB")
	}
	defer sess.Close()
	cur := sess.DB(db.Name()).C(coll)
	err = cur.Find(bson.M{"_id": id}).One(r)
	if err != nil {
		return r, errors.New("There was an error querying DB")
	}
	return r, nil
}

// New inserts a new Room in DB
func New(r *Room) (*Room, error) {
	db := config.DB{}
	sess, err := db.DoDial()
	if err != nil {
		return r, errors.New("There was an error connecting to Db")
	}
	defer sess.Close()
	r.ID = uuid.NewV1().String()
	r.RegistrationDate = time.Now()
	cur := sess.DB(db.Name()).C(coll)
	err = cur.Insert(r)
	if err != nil {
		return r, errors.New("There was an error creating a new Room")
	}
	return r, nil
}
