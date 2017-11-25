package room

import (
	"errors"
	"fmt"
	"time"

	"github.com/lalvarezguillen/roomies/config"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

//List fetches paginated rooms
func List(roomsQ *RoomsListQuery) RoomsLastID {
	db := config.DB{}
	sess, err := db.DoDial()
	if err != nil {
		panic("There was a problem connecting to the DB")
	}
	defer sess.Close()
	var rs Rooms
	coll := sess.DB(db.Name()).C(collection)
	var query bson.M
	if roomsQ.LastID != "" {
		query = bson.M{"_id": bson.M{"$lt": roomsQ.LastID}, "available": true}
	} else {
		query = bson.M{"available": true}
	}
	err = coll.Find(query).Sort("-ID").Limit(roomsQ.Limit).All(&rs)
	var lastID string
	if len(rs) > 0 {
		lastID = rs[len(rs)-1].ID
	}
	result := RoomsLastID{&rs, lastID}
	return result
}

// GetByID fetches a Room from DB by ID
func GetByID(id string) (*Room, error) {
	db := config.DB{}
	sess, err := db.DoDial()
	if err != nil {
		panic("There was a problem connecting to the DB")
	}
	defer sess.Close()
	r := Room{}
	coll := sess.DB(db.Name()).C(collection)
	err = coll.Find(bson.M{"_id": id}).One(&r)
	return &r, err
}

// New inserts a new Room in DB
func New(r *Room) (*Room, error) {
	db := config.DB{}
	sess, err := db.DoDial()
	if err != nil {
		panic("There was an error connecting to Db")
	}
	defer sess.Close()
	r.ID = uuid.NewV1().String()
	r.RegistrationDate = time.Now().UTC().Unix()
	r.Available = true
	coll := sess.DB(db.Name()).C(collection)
	err = coll.Insert(r)
	if err != nil {
		return r, errors.New("There was an error creating a new Room")
	}
	return r, nil
}

// Delete removes a room from DB
func Delete(id string) error {
	db := config.DB{}
	sess, err := db.DoDial()
	if err != nil {
		panic("There was a problem connecting to the DB")
	}
	defer sess.Close()
	coll := sess.DB(db.Name()).C(collection)
	return coll.RemoveId(id)
}

// Update updates a room's DB entry
func Update(id string, newData *map[string]interface{}) (*Room, error) {
	db := config.DB{}
	sess, err := db.DoDial()
	if err != nil {
		panic("There was a problem connecting to the DB")
	}
	defer sess.Close()
	coll := sess.DB(db.Name()).C(collection)
	updater := generateUpdater(newData)
	err = coll.UpdateId(id, &updater)
	if err != nil {
		return &Room{}, err
	}
	return GetByID(id)
}

func generateUpdater(newData *map[string]interface{}) bson.M {
	updates := bson.M{}
	for key, val := range *newData {
		updates[key] = val
	}
	updater := bson.M{"$set": updates}
	fmt.Println(updater)
	return updater
}
