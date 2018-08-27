package repository

func NotExists(field string) BinaryOperator {
	return &BinaryOperatorImpl{
		Type:      BinaryOperatorTypeExists,
		FieldName: field,
		Value:     false,
	}
}

func Like(field string, value interface{}) BinaryOperator {
	return &BinaryOperatorImpl{
		Type:      BinaryOperatorTypeLike,
		FieldName: field,
		Value:     value,
	}
}

func NotLike(field string, value interface{}) BinaryOperator {
	return &BinaryOperatorImpl{
		Type:      BinaryOperatorTypeNotLike,
		FieldName: field,
		Value:     value,
	}
}

func StartsWith(field string, value interface{}) BinaryOperator {
	return &BinaryOperatorImpl{
		Type:      BinaryOperatorTypeStartsWith,
		FieldName: field,
		Value:     value,
	}
}

func EndsWith(field string, value interface{}) BinaryOperator {
	return &BinaryOperatorImpl{
		Type:      BinaryOperatorTypeEndsWith,
		FieldName: field,
		Value:     value,
	}
}

func HasValue(field string) BinaryOperator {
	return &BinaryOperatorImpl{
		Type:      BinaryOperatorTypeHasValue,
		FieldName: field,
		Value:     true,
	}
}

func NotHasValue(field string) BinaryOperator {
	return &BinaryOperatorImpl{
		Type:      BinaryOperatorTypeNotHasValue,
		FieldName: field,
		Value:     false,
	}
}

func ElemMatch(attribute string, field string, value interface{}) BinaryOperator {
	return match(BinaryOperatorTypeElemMatch, attribute, field, "", value)
}

func NotElemMatch(attribute string, field string, value interface{}) BinaryOperator {
	return match(BinaryOperatorTypeNotElemMatch, attribute, field, "$ne", value)
}

func ElemMatchGT(attribute string, field string, value interface{}) BinaryOperator {
	return match(BinaryOperatorTypeNotElemMatch, attribute, field, "$gt", value)
}

func ElemMatchLT(attribute string, field string, value interface{}) BinaryOperator {
	return match(BinaryOperatorTypeNotElemMatch, attribute, field, "$lt", value)
}

func ElemMatchHasValue(attribute string, field string) BinaryOperator {
	return match(BinaryOperatorTypeNotElemMatch, attribute, field, "$exists", true)
}

func ElemMatchNotHasValue(attribute string, field string) BinaryOperator {
	return match(BinaryOperatorTypeNotElemMatch, attribute, field, "$exists", false)
}

func match(binaryType binaryOperatorType, attribute string, field string, op string, value interface{}) BinaryOperator {
	return &BinaryOperatorImpl{
		Type: binaryType,
		// OpField:   &op,
		// Attribute: &attribute,
		FieldName: field,
		Value:     value,
	}
}

func RelativeAfter(field string, value interface{}) BinaryOperator {
	return &BinaryOperatorImpl{
		Type:      BinaryOperatorRelativeAfter,
		FieldName: field,
		Value:     value,
	}
}

func RelativeBefore(field string, value interface{}) BinaryOperator {
	return &BinaryOperatorImpl{
		Type:      BinaryOperatorRelativeBefore,
		FieldName: field,
		Value:     value,
	}
}

func RelativeExactly(field string, value interface{}) BinaryOperator {
	return &BinaryOperatorImpl{
		Type:      BinaryOperatorRelativeExactly,
		FieldName: field,
		Value:     value,
	}
}

func Exactly(field string, value interface{}) BinaryOperator {
	return &BinaryOperatorImpl{
		Type:      BinaryOperatorExactly,
		FieldName: field,
		Value:     value,
	}
}

func Between(field string, value interface{}) BinaryOperator {
	return &BinaryOperatorImpl{
		Type:      BinaryOperatorBetween,
		FieldName: field,
		Value:     value,
	}
}
