package repository

func EQ(field string, value interface{}) BinaryOperator {
	return comparison(field, value, BinaryOperatorTypeEq, "$eq")
}

func GT(field string, value interface{}) BinaryOperator {
	return comparison(field, value, BinaryOperatorTypeGT, "$gt")
}

func GTE(field string, value interface{}) BinaryOperator {
	return comparison(field, value, BinaryOperatorTypeGTE, "$gte")
}

func IN(field string, value interface{}) BinaryOperator {
	return comparison(field, value, BinaryOperatorTypeIN, "$in")
}

func LT(field string, value interface{}) BinaryOperator {
	return comparison(field, value, BinaryOperatorTypeLT, "$lt")
}

func LTE(field string, value interface{}) BinaryOperator {
	return comparison(field, value, BinaryOperatorTypeLTE, "$lte")
}

func NE(field string, value interface{}) BinaryOperator {
	return comparison(field, value, BinaryOperatorTypeNE, "$ne")
}

func NIN(field string, value interface{}) BinaryOperator {
	return comparison(field, value, BinaryOperatorTypeNIN, "$nin")
}

func comparison(field string, value interface{}, operatorType binaryOperatorType, operator string) BinaryOperator {
	return &BinaryOperatorImpl{
		Type:      operatorType,
		OpField:   &operator,
		FieldName: field,
		Value:     value,
	}
}
