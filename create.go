package repository

import "github.com/globalsign/mgo"

func Create(r Repository, object interface{}) error {
	return r.GetQueryRunner().RunWithDB(func(db *mgo.Database) error {
		return db.C(r.GetCollectionName()).Insert(object)
	})
}
