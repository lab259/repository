package repository

import "github.com/globalsign/mgo"

func Delete(r Repository, params ...interface{}) error {
	return r.GetQueryRunner().RunWithDB(func(db *mgo.Database) error {
		c := db.C(r.GetCollectionName())
		criteria, err := GetQueryCriteria(r.GetDefaultCriteria(), params...)
		if err != nil {
			return err
		}
		return c.Remove(criteria)
	})
}
