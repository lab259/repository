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

func Or(conditions ... interface{}) *BooleanOperator {
	return &BooleanOperator{
		Type:       OperatorOr,
		Conditions: conditions,
	}
}

// TODO: interface conversion: interface {} is *repository.BinaryOperatorImpl, not repository.BinaryOperatorImpl
// func Not(conditions interface{}) BinaryOperator {
// 	cast := conditions.(*BinaryOperatorImpl)
// 	cast.Type = BinaryOperatorTypeNot
// 	return cast
// }

func Not(conditions interface{}) *BooleanOperator {
	// cond := []interface{}{conditions}
	b := conditions.(*BinaryOperatorImpl)
	cond := []interface{}{b}
	return &BooleanOperator{
		Field:      &b.FieldName,
		Type:       OperatorNot,
		Conditions: cond,
	}
}
