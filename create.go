package repository

import "gopkg.in/mgo.v2"

func Create(r Repository, object interface{}) error {
	return r.GetQueryRunner().RunWithDB(func(db *mgo.Database) error {
		return db.C(r.GetCollectionName()).Insert(object)
	})
}
