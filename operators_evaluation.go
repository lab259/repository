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

type FindText struct {
	Search             string  `bson:"$search"`
	Language           *string `bson:"$language,omitempty"`
	CaseSensitive      bool    `bson:"$caseSensitive,omitempty"`
	DiacriticSensitive bool    `bson:"$diacriticSensitive,omitempty"`
}

func Text(value ... interface{}) *BooleanOperator {
	return &BooleanOperator{
		Type:       OperatorText,
		Conditions: value,
	}
}
