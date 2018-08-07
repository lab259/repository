package repository

import "gopkg.in/mgo.v2"

func Delete(r Repository, params ... interface{}) error {
	return r.GetQueryRunner().RunWithDB(func(db *mgo.Database) error {
		c := db.C(r.GetCollectionName())
		criteria, err := GetQueryCriteria(r.GetDefaultCriteria(), params ...)
		if err != nil {
			return err
		}
		return c.Remove(criteria)
	})
}
