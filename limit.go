package repository

import (
	"github.com/globalsign/mgo"
)

type OperatorSkipLimit struct {
	Skip  int
	Limit int
}

func (o *OperatorSkipLimit) Apply(query *mgo.Query) (*mgo.Query, error) {
	if o.Skip > -1 {
		query = query.Skip(o.Skip)
	}
	if o.Limit > -1 {
		query = query.Limit(o.Limit)
	}
	return query, nil
}
