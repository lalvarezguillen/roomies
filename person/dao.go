package person

import (
	"errors"
	"time"

	"github.com/lalvarezguillen/roomies/config"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

const coll string = "people"

// GetByID obtains a Person by ID
func GetByID(id string) (Person, error) {
	db := config.DB{}
	p := Person{}
	sess, err := db.DoDial()
	if err != nil {
		return p, errors.New("There was a problem connecting to DB")
	}
	defer sess.Close()
	cur := sess.DB(db.Name()).C(coll)
	err = cur.Find(bson.M{"_id": id}).One(&p)
	if err != nil {
		return p, errors.New("There was an error querying DB")
	}
	return p, nil
}

func New(p Person) (Person, error) {
	db := config.DB{}
	sess, err := db.DoDial()
	if err != nil {
		return p, errors.New("There was an error connecting to DB")
	}
	defer sess.Close()
	p.ID = uuid.NewV1().String()
	p.RegistrationDate = time.Now()
	cur := sess.DB(db.Name()).C(coll)
	err = cur.Insert(p)
	if err != nil {
		return p, errors.New("There was an error creatting a new Person")
	}
	return p, nil
}
