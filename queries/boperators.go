package queries

import ".."

func EQ(field string, value interface{}) repository.BinaryOperator {
	return &repository.BinaryOperatorImpl{
		Type:      repository.BinaryOperatorTypeEq,
		FieldName: field,
		Value:     value,
	}
}

func GT(field string, value interface{}) repository.BinaryOperator {
	return &repository.BinaryOperatorImpl{
		Type:      repository.BinaryoperatorTypeGT,
		FieldName: field,
		Value:     value,
	}
}

func GTE(field string, value interface{}) repository.BinaryOperator {
	return &repository.BinaryOperatorImpl{
		Type:      repository.BinaryoperatorTypeGTE,
		FieldName: field,
		Value:     value,
	}
}

func LT(field string, value interface{}) repository.BinaryOperator {
	return &repository.BinaryOperatorImpl{
		Type:      repository.BinaryoperatorTypeLT,
		FieldName: field,
		Value:     value,
	}
}

func LTE(field string, value interface{}) repository.BinaryOperator {
	return &repository.BinaryOperatorImpl{
		Type:      repository.BinaryoperatorTypeLTE,
		FieldName: field,
		Value:     value,
	}
}

func Exists(field string) repository.BinaryOperator {
	return &repository.BinaryOperatorImpl{
		Type:      repository.BinaryoperatorTypeExists,
		FieldName: field,
		Value:     true,
	}
}

func NotExists(field string) repository.BinaryOperator {
	return &repository.BinaryOperatorImpl{
		Type:      repository.BinaryoperatorTypeExists,
		FieldName: field,
		Value:     false,
	}
}
