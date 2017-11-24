package room

import (
	"errors"
	"time"

	"github.com/lalvarezguillen/roomies/config"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

//ListRooms fetches paginated rooms
func ListRooms(roomsQ *RoomsListQuery) (RoomsLastID, error) {
	db := config.DB{}
	sess, err := db.DoDial()
	var rs Rooms
	result := RoomsLastID{&rs, ""}
	if err != nil {
		return result, errors.New("There was a problem connecting to the DB")
	}
	defer sess.Close()
	cur := sess.DB(db.Name()).C(coll)
	var query bson.M
	if roomsQ.LastID != "" {
		query = bson.M{"_id": bson.M{"$lt": roomsQ.LastID}, "available": true}
	} else {
		query = bson.M{"available": true}
	}
	err = cur.Find(query).Sort("-ID").Limit(roomsQ.Limit).All(rs)
	result = RoomsLastID{&rs, rs[len(rs)-1].ID}
	return result, nil
}

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
