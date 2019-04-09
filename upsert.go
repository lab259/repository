package repository

import (
	"github.com/globalsign/mgo"
)

// Upsert find a single document, if matching update else create.
func Upsert(r Repository, object interface{}, selector ...interface{}) error {
	return r.GetQueryRunner().RunWithDB(func(db *mgo.Database) error {
		criteria, err := GetQueryCriteria(r.GetDefaultCriteria(), selector...)
		if err != nil {
			return err
		}

		_, err = db.C(r.GetCollectionName()).Upsert(criteria, object)
		if err != nil {
			return err
		}

		return err
	})
}
