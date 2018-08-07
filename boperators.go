package repository

func EQ(field string, value interface{}) BinaryOperator {
	return &BinaryOperatorImpl{
		Type:      BinaryOperatorTypeEq,
		FieldName: field,
		Value:     value,
	}
}

func GT(field string, value interface{}) BinaryOperator {
	return &BinaryOperatorImpl{
		Type:      BinaryoperatorTypeGT,
		FieldName: field,
		Value:     value,
	}
}

func GTE(field string, value interface{}) BinaryOperator {
	return &BinaryOperatorImpl{
		Type:      BinaryoperatorTypeGTE,
		FieldName: field,
		Value:     value,
	}
}

func LT(field string, value interface{}) BinaryOperator {
	return &BinaryOperatorImpl{
		Type:      BinaryoperatorTypeLT,
		FieldName: field,
		Value:     value,
	}
}

func LTE(field string, value interface{}) BinaryOperator {
	return &BinaryOperatorImpl{
		Type:      BinaryoperatorTypeLTE,
		FieldName: field,
		Value:     value,
	}
}

func Exists(field string) BinaryOperator {
	return &BinaryOperatorImpl{
		Type:      BinaryoperatorTypeExists,
		FieldName: field,
		Value:     true,
	}
}

func NotExists(field string) BinaryOperator {
	return &BinaryOperatorImpl{
		Type:      BinaryoperatorTypeExists,
		FieldName: field,
		Value:     false,
	}
}
