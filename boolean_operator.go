package repository

import (
	"github.com/globalsign/mgo/bson"
)

type operatorType int

const (
	OperatorAnd operatorType = iota
	OperatorNot
	OperatorNor
	OperatorOr
)

type BooleanOperator struct {
	Field      *string
	Type       operatorType
	Conditions []interface{}
}

func (o *BooleanOperator) GetCondition() (bson.DocElem, error) {
	var t string
	switch o.Type {
	case OperatorAnd:
		t = "$and"
	case OperatorNor:
		t = "$nor"
	case OperatorOr:
		t = "$or"
	case OperatorNot:
		t = "$not"
		cast := *o.Conditions[0].(*BinaryOperatorImpl)
		return bson.DocElem{
			Name: *o.Field,
			Value: bson.M{
				t: bson.M{
					*cast.OpField: cast.Value,
				},
			},
		}, nil
	}
	conds := make([]interface{}, 0, len(o.Conditions))
	for _, cond := range o.Conditions {
		switch condition := cond.(type) {
		case QueryBuilder:
			ccc, err := condition.GetQuery()
			if err != nil {
				return bson.DocElem{}, err
			}
			if ccc != nil {
				conds = append(conds, ccc)
			}
		case BinaryOperator:
			ccc, err := condition.GetCondition()
			if err != nil {
				return bson.DocElem{}, err
			}
			conds = append(conds, bson.D{ccc})
		}
	}
	return bson.DocElem{
		Name:  t,
		Value: conds,
	}, nil
}
