package queries

import ".."

func WithPage(page, pageSize int) *repository.OperatorSkipLimit {
	return &repository.OperatorSkipLimit{Skip: page * pageSize, Limit: pageSize}
}

func Skip(skip int) *repository.OperatorSkipLimit {
	return &repository.OperatorSkipLimit{Skip: skip, Limit: -1}
}

func Limit(limit int) *repository.OperatorSkipLimit {
	return &repository.OperatorSkipLimit{Skip: -1, Limit: limit}
}

func ByID(id interface{}) repository.BinaryOperator {
	return &repository.BinaryOperatorImpl{
		FieldName: "_id",
		Type:      repository.BinaryOperatorTypeEq,
		Value:     id,
	}
}

func WithCriteria(params ... interface{}) *repository.Criteria {
	return &repository.Criteria{
		Conditions: params,
	}
}

func Sort(fields ... string) *repository.Sort {
	return &repository.Sort{
		Fields: fields,
	}
}

func And(conditions ... interface{}) *repository.BooleanOperator {
	return &repository.BooleanOperator{
		Type:       repository.OperatorAnd,
		Conditions: conditions,
	}
}

func Or(conditions ... interface{}) *repository.BooleanOperator {
	return &repository.BooleanOperator{
		Type:       repository.OperatorOr,
		Conditions: conditions,
	}
}
