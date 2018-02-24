package person

import (
	"errors"
	"time"

	"github.com/lalvarezguillen/roomies/config"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

//List fetches paginated people
func List(peopleQ *PeopleListQuery) PeopleQueryResult {
	db := config.DB{}
	sess, err := db.DoDial()
	if err != nil {
		panic("The was a problem connecting to DB")
	}
	defer sess.Close()
	var ps People
	coll := sess.DB(db.Name()).C(Collection)
	var query bson.M
	if peopleQ.LastID != "" {
		query = bson.M{"_id": bson.M{"$lt": peopleQ.LastID}}
	}
	err = coll.Find(query).Sort("-ID").Limit(peopleQ.Limit).All(&ps)
	var lastID string
	if len(ps) > 0 {
		lastID = ps[len(ps)-1].ID
	}
	result := PeopleQueryResult{&ps, lastID}
	return result
}

// GetByID obtains a Person by ID
func GetByID(id string) (*Person, error) {
	db := config.DB{}
	p := &Person{}
	sess, err := db.DoDial()
	if err != nil {
		return p, errors.New("There was a problem connecting to DB")
	}
	defer sess.Close()
	cur := sess.DB(db.Name()).C(Collection)
	err = cur.Find(bson.M{"_id": id}).One(p)
	if err != nil {
		return p, errors.New("There was an error querying DB")
	}
	return p, nil
}

// New creates a new Person
func New(p *Person) (*Person, error) {
	db := config.DB{}
	sess, err := db.DoDial()
	if err != nil {
		return p, errors.New("There was an error connecting to DB")
	}
	defer sess.Close()
	p.ID = uuid.NewV1().String()
	p.RegistrationDate = time.Now()
	cur := sess.DB(db.Name()).C(Collection)
	err = cur.Insert(p)
	if err != nil {
		return p, errors.New("There was an error creatting a new Person")
	}
	return p, nil
}

// Delete removes a Person
func Delete(id string) error {
	db := config.DB{}
	sess, err := db.DoDial()
	if err != nil {
		panic("There was a problem connecting to the DB")
	}
	defer sess.Close()
	coll := sess.DB(db.Name()).C(Collection)
	return coll.RemoveId(id)
}

// Update updates a Person's DB entry
func Update(p *Person) (*Person, error) {
	db := config.DB{}
	sess, err := db.DoDial()
	if err != nil {
		panic("There was a problem connecting to the DB")
	}
	defer sess.Close()
	coll := sess.DB(db.Name()).C(Collection)
	err = coll.UpdateId(&p.ID, &p)
	return p, err
}
