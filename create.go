package repository

import "github.com/globalsign/mgo"

func Create(r Repository, object interface{}) error {
	return r.GetQueryRunner().RunWithDB(func(db *mgo.Database) error {
		return db.C(r.GetCollectionName()).Insert(object)
	})
}

func CreateIndexes(r Repository, indexes ...interface{}) error {
	return r.GetQueryRunner().RunWithDB(func(db *mgo.Database) error {
		var err error
		for _, i := range indexes {
			err = db.C(r.GetCollectionName()).EnsureIndex(i.(mgo.Index))
			if err != nil {
				return err
			}
		}
		return err
	})
}