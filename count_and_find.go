package repository

import (
	"github.com/globalsign/mgo"
)

func CountAndFindAll(r Repository, dst interface{}, params ...interface{}) (n int, err error) {
	err = r.GetQueryRunner().RunWithDB(func(db *mgo.Database) error {
		query, count, err := CountAndQuery(r, db.C(r.GetCollectionName()), params...)
		if err != nil {
			return err
		}
		err = query.All(dst)
		if err != nil {
			return err
		}
		n = count
		return nil
	})
	return
}
