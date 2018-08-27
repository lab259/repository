package repository

import (
	"github.com/globalsign/mgo/bson"
)

type binaryOperatorType int

const (
	// Comparison Operators
	BinaryOperatorTypeEq = iota
	BinaryOperatorTypeGT
	BinaryOperatorTypeGTE
	BinaryOperatorTypeIN
	BinaryOperatorTypeLT
	BinaryOperatorTypeLTE
	BinaryOperatorTypeNE
	BinaryOperatorTypeNIN

	// Element Operators
	BinaryOperatorTypeExists

	// Evaluation Operators
	BinaryOperatorTypeRegex
)

type BinaryOperator interface {
	GetCondition() (bson.DocElem, error)
}

type BinaryOperatorImpl struct {
	Attribute *string
	OpField   *string
	FieldName string
	Type      binaryOperatorType
	Value     interface{}
}

func (o *BinaryOperatorImpl) GetCondition() (bson.DocElem, error) {
	switch o.Type {
	case BinaryOperatorTypeEq:
		return bson.DocElem{Name: o.FieldName, Value: o.Value}, nil
	}
	return bson.DocElem{
		Name: o.FieldName,
		Value: bson.M{
			*o.OpField: o.Value,
		},
	}, nil
}
