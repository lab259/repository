package repository

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Update does a complete update of a model.
func Update(r Repository, id interface{}, object interface{}) error {
	return r.GetQueryRunner().RunWithDB(func(db *mgo.Database) error {
		return db.C(r.GetCollectionName()).Update(bson.D{
			bson.DocElem{"_id", id}}, bson.M{
			"$set": object,
		})
	})
}
