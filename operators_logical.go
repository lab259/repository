package repository

func And(conditions ... interface{}) *BooleanOperator {
	return &BooleanOperator{
		Type:       OperatorAnd,
		Conditions: conditions,
	}
}

func Nor(conditions ... interface{}) *BooleanOperator {
	return &BooleanOperator{
		Type:       OperatorNor,
		Conditions: conditions,
	}
}

func Not(conditions interface{}) *BooleanOperator {
	b := conditions.(*BinaryOperatorImpl)
	cond := []interface{}{b}
	return &BooleanOperator{
		Field:      &b.FieldName,
		Type:       OperatorNot,
		Conditions: cond,
	}
}

func Or(conditions ... interface{}) *BooleanOperator {
	return &BooleanOperator{
		Type:       OperatorOr,
		Conditions: conditions,
	}
}