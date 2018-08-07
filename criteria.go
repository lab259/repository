package repository

import (
		"gopkg.in/mgo.v2/bson"
)

type Criteria struct {
	Conditions []interface{}
}

func (c *Criteria) GetQuery() (bson.D, error) {
	result := make(bson.D, 0, len(c.Conditions))
	for _, cond := range c.Conditions {
		switch condition := cond.(type) {
		case bson.DocElem:
			result = append(result, condition)
		case BinaryOperator:
			ccc, err := condition.GetCondition()
			if err != nil {
				return nil, err
			}
			result = append(result, ccc)
		default:
			return nil, NewErrTypeNotSupported(cond)
		}
	}
	return result, nil
}

type RawCriteria bson.D

func (c *RawCriteria) GetQuery() (bson.D, error) {
	return bson.D(*c), nil
}
