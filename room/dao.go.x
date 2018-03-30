package room

import (
	"errors"
	"time"

	"github.com/lalvarezguillen/roomies/config"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

//ListRooms fetches paginated rooms
func ListRooms(roomsQ *RoomsListQuery) RoomsQueryResult {
	db := config.DB{}
	sess, err := db.DoDial()
	if err != nil {
		panic("There was a problem connecting to the DB")
	}
	defer sess.Close()
	var rs Rooms
	coll := sess.DB(db.Name()).C(Collection)
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
	result := RoomsQueryResult{&rs, lastID}
	return result
}

// GetRoomByID fetches a Room from DB by ID
func GetRoomByID(id string) (*Room, error) {
	db := config.DB{}
	sess, err := db.DoDial()
	if err != nil {
		panic("There was a problem connecting to the DB")
	}
	defer sess.Close()
	r := Room{}
	coll := sess.DB(db.Name()).C(Collection)
	err = coll.Find(bson.M{"_id": id}).One(&r)
	return &r, err
}

// NewRoom inserts a new Room in DB
func NewRoom(r *Room) (*Room, error) {
	db := config.DB{}
	sess, err := db.DoDial()
	if err != nil {
		panic("There was an error connecting to Db")
	}
	defer sess.Close()
	u, err := uuid.NewV1()
	if err != nil {
		panic(err)
	}
	r.ID = u.String()
	r.RegistrationDate = time.Now().UTC().Unix()
	r.Available = true
	coll := sess.DB(db.Name()).C(Collection)
	err = coll.Insert(r)
	if err != nil {
		return r, errors.New("There was an error creating a new Room")
	}
	return r, nil
}

// DeleteRoom removes a room from DB
func DeleteRoom(id string) error {
	db := config.DB{}
	sess, err := db.DoDial()
	if err != nil {
		panic("There was a problem connecting to the DB")
	}
	defer sess.Close()
	coll := sess.DB(db.Name()).C(Collection)
	return coll.RemoveId(id)
}

// UpdateRoom updates a room's DB entry
func UpdateRoom(r *Room) (*Room, error) {
	db := config.DB{}
	sess, err := db.DoDial()
	if err != nil {
		panic("There was a problem connecting to the DB")
	}
	defer sess.Close()
	coll := sess.DB(db.Name()).C(Collection)
	err = coll.UpdateId(&r.ID, &r)
	return r, err
}
