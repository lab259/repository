package repository

import (
	"gopkg.in/mgo.v2"
)

// FindByID finds a model based on its ID. If not model is found in the
// collection, this method will return a mgo.ErrNotFound error.
//
// When generating the query for the MongoDB, this method will use the
// defaultCriteria.
//
// If no model is found, it returns an `mgo.ErrNotFound`.
func FindByID(r Repository, id interface{}, dst interface{}, params ... interface{}) error {
	return r.GetQueryRunner().RunWithDB(func(db *mgo.Database) error {
		query, err := Query(r, db.C(r.GetCollectionName()), append([]interface{}{
			&Criteria{
				[]interface{}{
					&BinaryOperatorImpl{
						Type:      BinaryOperatorTypeEq,
						FieldName: "_id",
						Value:     id,
					},
				},
			},
		}, params...)...)
		if err != nil {
			return err
		}
		return query.Limit(1).One(dst)
	})
}
