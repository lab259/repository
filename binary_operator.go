package repository

import (
	"gopkg.in/mgo.v2/bson"
)

type binaryOperatorType int

const (
	BinaryOperatorTypeEq = iota
	BinaryoperatorTypeGT
	BinaryoperatorTypeLT
	BinaryoperatorTypeGTE
	BinaryoperatorTypeLTE
	BinaryoperatorTypeExists
	BinaryoperatorTypeNotExists
)

type BinaryOperator interface {
	GetCondition() (bson.DocElem, error)
}

type BinaryOperatorImpl struct {
	FieldName string
	Type      binaryOperatorType
	Value     interface{}
}

func (o *BinaryOperatorImpl) GetCondition() (bson.DocElem, error) {
	var operator string
	switch o.Type {
	case BinaryOperatorTypeEq:
		return bson.DocElem{Name: o.FieldName, Value: o.Value}, nil
	case BinaryoperatorTypeGT:
		operator = "$gt"
	case BinaryoperatorTypeGTE:
		operator = "$gte"
	case BinaryoperatorTypeLT:
		operator = "$lt"
	case BinaryoperatorTypeLTE:
		operator = "$lte"
	case BinaryoperatorTypeExists:
		operator = "$exists"
	}
	return bson.DocElem{
		Name: o.FieldName,
		Value: bson.M{
			operator: o.Value,
		},
	}, nil
}
