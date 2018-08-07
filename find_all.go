package repository

import (
	"gopkg.in/mgo.v2"
)

func FindAll(r Repository, objects interface{}, params ... interface{}) error {
	return r.GetQueryRunner().RunWithDB(func(db *mgo.Database) error {
		query, err := Query(r, db.C(r.GetCollectionName()), params...)
		if err != nil {
			return err
		}
		return query.All(objects)
	})
}
