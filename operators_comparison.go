package repository

func EQ(field string, value interface{}) BinaryOperator {
	return comparison(field, value, BinaryOperatorTypeEq)
}

func GT(field string, value interface{}) BinaryOperator {
	return comparison(field, value, BinaryOperatorTypeGT)
}

func GTE(field string, value interface{}) BinaryOperator {
	return comparison(field, value, BinaryOperatorTypeGTE)
}

func IN(field string, value interface{}) BinaryOperator {
	return comparison(field, value, BinaryOperatorTypeIN)
}

func LT(field string, value interface{}) BinaryOperator {
	return comparison(field, value, BinaryOperatorTypeLT)
}

func LTE(field string, value interface{}) BinaryOperator {
	return comparison(field, value, BinaryOperatorTypeLTE)
}

func NE(field string, value interface{}) BinaryOperator {
	return comparison(field, value, BinaryOperatorTypeNE)
}

func NIN(field string, value interface{}) BinaryOperator {
	return comparison(field, value, BinaryOperatorTypeNIN)
}

func comparison(field string, value interface{}, operator binaryOperatorType) BinaryOperator {
	return &BinaryOperatorImpl{
		Type:      operator,
		FieldName: field,
		Value:     value,
	}
}
