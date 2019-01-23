package repository

func Exists(field string, value bool) BinaryOperator {
	operator := "$exists"
	return &BinaryOperatorImpl{
		Type:      BinaryOperatorTypeExists,
		OpField:   &operator,
		FieldName: field,
		Value:     value,
	}
}
