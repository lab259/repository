package repository

import (
	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2/bson"
)

type projection struct {
	fields []interface{}
}

func (f *projection) Apply(query *mgo.Query) (*mgo.Query, error) {
	fields := make(bson.M, len(f.fields))
	for _, mfield := range f.fields {
		switch field := mfield.(type) {
		case string:
			fname := field
			fvalue := 1
			switch field[0] {
			case '+':
				fname = fname[1:]
			case '-':
				fname = fname[1:]
				fvalue = 0
			}
			fields[fname] = fvalue
		case bson.DocElem:
			fields[field.Name] = field.Value
		case *bson.DocElem:
			fields[field.Name] = field.Value
		default:
			return nil, NewErrTypeNotSupported(mfield)
		}
	}
	return query.Select(fields), nil
}

func Projection(fields ...interface{}) *projection {
	return &projection{
		fields: fields,
	}
}
