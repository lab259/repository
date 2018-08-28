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

	// Evaluation Operators
	OperatorText

	// Array Operators
	OperatorTypeElemMatch
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
	case OperatorNot: // Consider behaviors using the $not and $regex
		t = "$not"
		cast := *o.Conditions[0].(*BinaryOperatorImpl)
		if cast.Type == BinaryOperatorTypeRegex {
			return bson.DocElem{
				Name: *o.Field,
				Value: bson.M{
					t: cast.Value,
				},
			}, nil
		} else {
			return bson.DocElem{
				Name: *o.Field,
				Value: bson.M{
					t: bson.M{
						*cast.OpField: cast.Value,
					},
				},
			}, nil
		}
	case OperatorText:
		t = "$text"
		return bson.DocElem{
			Name:  t,
			Value: o.Conditions[0].(FindText),
		}, nil
	case OperatorTypeElemMatch:
		t = "$elemMatch"
		elem := make(bson.D, 0, len(o.Conditions))
		for i, cond := range o.Conditions {
			switch cond.(type) {
			case *BooleanOperator:
				p, err := o.Conditions[i].(*BooleanOperator).GetCondition()
				if err != nil {
					return bson.DocElem{}, NewErrTypeNotSupported(cond)
				}
				elem = append(elem, p)
			case BinaryOperator:
				p, err := o.Conditions[i].(*BinaryOperatorImpl).GetCondition()
				if err != nil {
					return bson.DocElem{}, NewErrTypeNotSupported(cond)
				}
				elem = append(elem, p)
			default:
				return bson.DocElem{}, NewErrTypeNotSupported(cond)
			}
		}
		return bson.DocElem{
			Name: *o.Field,
			Value: bson.M{
				t: elem,
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
