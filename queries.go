package repository

func WithPage(page, pageSize int) *OperatorSkipLimit {
	return &OperatorSkipLimit{Skip: page * pageSize, Limit: pageSize}
}

func Skip(skip int) *OperatorSkipLimit {
	return &OperatorSkipLimit{Skip: skip, Limit: -1}
}

func Limit(limit int) *OperatorSkipLimit {
	return &OperatorSkipLimit{Skip: -1, Limit: limit}
}

func ByID(id interface{}) BinaryOperator {
	return &BinaryOperatorImpl{
		FieldName: "_id",
		Type:      BinaryOperatorTypeEq,
		Value:     id,
	}
}

func WithCriteria(params ... interface{}) *Criteria {
	return &Criteria{
		Conditions: params,
	}
}

func WithSort(fields ... string) *Sort {
	return &Sort{
		Fields: fields,
	}
}

func And(conditions ... interface{}) *BooleanOperator {
	return &BooleanOperator{
		Type:       OperatorAnd,
		Conditions: conditions,
	}
}

func Or(conditions ... interface{}) *BooleanOperator {
	return &BooleanOperator{
		Type:       OperatorOr,
		Conditions: conditions,
	}
}
