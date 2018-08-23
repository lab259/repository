package repository

import (
	"github.com/globalsign/mgo/bson"
)

type binaryOperatorType int

const (
	// Comparison Operators
	BinaryOperatorTypeEq  = iota
	BinaryOperatorTypeGT
	BinaryOperatorTypeGTE
	BinaryOperatorTypeIN
	BinaryOperatorTypeLT
	BinaryOperatorTypeLTE
	BinaryOperatorTypeNE
	BinaryOperatorTypeNIN

	// Operators
	BinaryOperatorTypeExists
	BinaryOperatorTypeLike
	BinaryOperatorTypeNotLike
	BinaryOperatorTypeStartsWith
	BinaryOperatorTypeEndsWith
	BinaryOperatorTypeHasValue
	BinaryOperatorTypeNotHasValue
	BinaryOperatorTypeElemMatch
	BinaryOperatorTypeNotElemMatch
	BinaryOperatorRelativeAfter
	BinaryOperatorRelativeBefore
	BinaryOperatorRelativeExactly
	BinaryOperatorExactly
	BinaryOperatorBetween
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
	var operator string
	switch o.Type {
	case BinaryOperatorTypeEq:
		return bson.DocElem{Name: o.FieldName, Value: o.Value}, nil
	case BinaryOperatorTypeGT:
		operator = "$gt"
	case BinaryOperatorTypeGTE:
		operator = "$gte"
	case BinaryOperatorTypeIN:
		operator = "$in"
	case BinaryOperatorTypeLT:
		operator = "$lt"
		o.OpField = &operator
	case BinaryOperatorTypeLTE:
		operator = "$lte"
	case BinaryOperatorTypeNIN:
		operator = "$nin"
	case BinaryOperatorTypeNE:
		operator = "$ne"
	case BinaryOperatorTypeExists:
		operator = "$exists"
		/*
		case BinaryOperatorTypeLike:
			return bson.DocElem{
				Name:  o.FieldName,
				Value: bson.RegEx{Pattern: o.Value.(string), Options: "i"},
			}, nil
		case BinaryOperatorTypeNotLike:
			operator = "$not"
			return bson.DocElem{
				Name: o.FieldName,
				Value: bson.M{
					operator: bson.RegEx{Pattern: o.Value.(string), Options: "i"},
				},
			}, nil
		case BinaryOperatorTypeStartsWith:
			return bson.DocElem{
				Name:  o.FieldName,
				Value: bson.RegEx{Pattern: "^" + o.Value.(string), Options: "i"},
			}, nil
		case BinaryOperatorTypeEndsWith:
			return bson.DocElem{
				Name:  o.FieldName,
				Value: bson.RegEx{Pattern: o.Value.(string) + "$", Options: "i"},
			}, nil
		case BinaryOperatorTypeHasValue:
			operator = "$exists"
			return bson.DocElem{
				Name: o.FieldName,
				Value: bson.M{
					operator: true,
				},
			}, nil
		case BinaryOperatorTypeNotHasValue:
			operator = "$exists"
			return bson.DocElem{
				Name: o.FieldName,
				Value: bson.M{
					operator: false,
				},
			}, nil
		case BinaryOperatorTypeElemMatch:
			operator = "$elemMatch"
			return bson.DocElem{
				Name: *o.Attribute,
				Value: bson.M{
					operator: bson.M{
						"name":  o.FieldName,
						"value": o.Value,
					},
				},
			}, nil
		case BinaryOperatorTypeNotElemMatch:
			operator = "$elemMatch"
			return bson.DocElem{
				Name: *o.Attribute,
				Value: bson.M{
					operator: bson.M{
						"name": o.FieldName,
						"value": bson.M{
							*o.OpField: o.Value,
						},
					},
				},
			}, nil
		case BinaryOperatorRelativeAfter:
			operator = "$gt"
			now := time.Now().UTC()
			days, err := strconv.Atoi(o.Value.(string))
			if err != nil {
				return bson.DocElem{}, err
			}
			sub := now.Add(-time.Duration(days) * time.Hour * 24)
			return bson.DocElem{
				Name: o.FieldName,
				Value: bson.M{
					operator: sub,
				},
			}, nil
		case BinaryOperatorRelativeBefore:
			operator = "$lt"
			now := time.Now().UTC()
			days, err := strconv.Atoi(o.Value.(string))
			if err != nil {
				return bson.DocElem{}, err
			}
			sub := now.Add(-time.Duration(days) * time.Hour * 24)
			return bson.DocElem{
				Name: o.FieldName,
				Value: bson.M{
					operator: sub,
				},
			}, nil
		case BinaryOperatorRelativeExactly:
			now := time.Now().UTC()
			day, err := strconv.Atoi(o.Value.(string))
			if err != nil {
				return bson.DocElem{}, err
			}
			starts := now.Add(-time.Duration(day) * time.Hour * 24)
			starts = time.Date(starts.Year(), starts.Month(), starts.Day(), 0, 0, 0, 0, time.UTC)
			ends := now.Add(-time.Duration(day) * time.Hour * 24)
			ends = time.Date(ends.Year(), ends.Month(), ends.Day(), 23, 59, 59, 0, time.UTC)

			return bson.DocElem{
				Name: o.FieldName,
				Value: map[string]interface{}{
					"$gte": starts,
					"$lte": ends,
				},
			}, nil
		case BinaryOperatorExactly:
			date := o.Value.(time.Time)
			starts := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
			ends := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, time.UTC)

			return bson.DocElem{
				Name: o.FieldName,
				Value: map[string]interface{}{
					"$gte": starts,
					"$lte": ends,
				},
			}, nil
		case BinaryOperatorBetween:
			fmt.Println("Between")
			between := o.Value.([]*string)
			starts, err := time.Parse(time.RFC3339, *between[0])
			if err != nil {
				return bson.DocElem{}, err
			}
			ends, err := time.Parse(time.RFC3339, *between[1])
			if err != nil {
				return bson.DocElem{}, err
			}

			return bson.DocElem{
				Name: o.FieldName,
				Value: map[string]interface{}{
					"$gte": starts,
					"$lte": ends,
				},
			}, nil
		*/
	}
	return bson.DocElem{
		Name: o.FieldName,
		Value: bson.M{
			operator: o.Value,
		},
	}, nil
}
