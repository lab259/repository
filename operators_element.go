package repository

func Exists(field string, value interface{}) BinaryOperator {
	operator := "$exists"
	return &BinaryOperatorImpl{
		Type:      BinaryOperatorTypeExists,
		OpField:   &operator,
		FieldName: field,
		Value:     value,
	}
}

// TODO: implements operator type. See: https://docs.mongodb.com/manual/reference/operator/query/type/
func Type(field string, value interface{}) BinaryOperator {
	operator := "$type"
	return &BinaryOperatorImpl{
		Type:      BinaryOperatorType,
		OpField:   &operator,
		FieldName: field,
		Value:     value,
	}
}