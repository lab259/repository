package repository

import "github.com/globalsign/mgo/bson"

func Regex(field string, value interface{}, options string) BinaryOperator {
	operator := "$regex"
	if options == "" {
		options = "i"
	}
	return &BinaryOperatorImpl{
		Type:      BinaryOperatorTypeRegex,
		OpField:   &operator,
		FieldName: field,
		Value: bson.RegEx{
			Pattern: value.(string),
			Options: options,
		},
	}
}

func Text(field string, value interface{}) BinaryOperator {
	operator := "$text"
	return &BinaryOperatorImpl{
		Type:      BinaryOperatorTypeText,
		OpField:   &operator,
		FieldName: field,
		Value:     value,
	}
}
