package repository

import "gopkg.in/mgo.v2"

func DeleteAll(r Repository, params ... interface{}) (n int, resultErr error) {
	resultErr = r.GetQueryRunner().RunWithDB(func(db *mgo.Database) error {
		c := db.C(r.GetCollectionName())
		criteria, err := GetQueryCriteria(r.GetDefaultCriteria(), params...)
		if err != nil {
			return err
		}
		info, err := c.RemoveAll(criteria)
		if err != nil {
			return err
		}
		n = info.Removed
		return nil
	})
	return
}
