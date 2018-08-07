package repository

import (
	"gopkg.in/mgo.v2/bson"
)

type operatorType int

const (
	OperatorAnd operatorType = iota
	OperatorOr
)

type BooleanOperator struct {
	Type       operatorType
	Conditions []interface{}
}

func (o *BooleanOperator) GetCondition() (bson.DocElem, error) {
	var t string
	switch o.Type {
	case OperatorAnd:
		t = "$and"
	case OperatorOr:
		t = "$or"
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
