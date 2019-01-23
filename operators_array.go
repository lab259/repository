package repository

func ElemMatch(field string, value ...interface{}) *BooleanOperator {
	return &BooleanOperator{
		Field:      &field,
		Type:       OperatorTypeElemMatch,
		Conditions: value,
	}
}
