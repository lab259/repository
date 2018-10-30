package repository

import (
	"github.com/globalsign/mgo"
)

// UpdateRawWithCriteria does a update of a model based on any selector.
func UpdateRawWithCriteria(r Repository, object interface{}, selector ...interface{}) error {
	return r.GetQueryRunner().RunWithDB(func(db *mgo.Database) error {
		criteria, err := GetQueryCriteria(r.GetDefaultCriteria(), selector...)
		if err != nil {
			return err
		}
		info, err := db.C(r.GetCollectionName()).UpdateAll(criteria, object)
		if err != nil {
			return err
		}

		if info.Matched == 0 {
			return mgo.ErrNotFound
		}
		return err
	})
}
