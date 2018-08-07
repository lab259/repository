package repository

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
		"reflect"
)

// Queryable objects will be treated as special objects during the query
// building.
type QueryBuilder interface {
	// `GetQuery` is used for filtering information. It is the object that
	// would be passed to the `mgo.Collection.Find` method to return a
	// `mgo.Query`. But it returns just a piece of it for further aggregation
	// inside of the `Query` method.
	GetQuery() (bson.D, error)
}

// QueryModifier is the interface that describes an object that can modify one
// `mgo.Query`
type QueryModifier interface {
	// Apply will
	Apply(query *mgo.Query) (*mgo.Query, error)
}

func genConditions(conditions bson.D, params ... interface{}) (bson.D, error) {
	if conditions == nil {
		conditions = make(bson.D, 0, len(params))
	}
	for _, param := range params {
		switch v := param.(type) {
		case bson.D:
			// Appends the conditions directly to the result.
			conditions = append(conditions, v...)
			break
		case BinaryOperator:
			// If a `BinaryOperator` is passed.
			// Get the conditions
			cond, err := v.GetCondition()
			if err != nil {
				return nil, err
			}
			// Applies it to the result
			conditions = append(conditions, cond)
		case QueryBuilder:
			cond, err := v.GetQuery()
			if err != nil {
				return nil, err
			}
			if cond != nil {
				conditions = append(conditions, cond...)
			}
		}
	}
	return conditions, nil
}

// GetQueryCriteria builds the criteria that will be performed in the Find
// method of a `mgo.Collection.Find` method.
//
// The parameter `defaultCriteria` represents an initial set of criterias that
// be mixed with the passed parameters.
//
// An error can be returned when building the criteria, due to a non supported
// datatype being used. Otherwise, a ready for use criteria object is returned.
func GetQueryCriteria(defaultCriteria interface{}, params ... interface{}) (bson.D, error) {
	var conditions bson.D
	if defaultCriteria == nil {
		conditions = make(bson.D, 0, len(params))
	} else {
		ccc, err := genConditions(conditions, defaultCriteria)
		if err != nil {
			return nil, err
		}
		conditions = ccc
	}
	ccc, err := genConditions(conditions, params...)
	if err != nil {
		return nil, err
	}
	return ccc, nil
}

// Query builds the criteria, using `GetQueryCriteria`, and performs the
// `mgo.Collection.Find` for building the `mgo.Query`.
//
// Then, this generated `mgo.Query` will passed for the params that are from the
// `Queryable` type, applying its transformations (through `Queryable.Applying`).
//
// Finally, it will return
func Query(r Repository, c *mgo.Collection, params ...interface{}) (*mgo.Query, error) {
	conditions, err := GetQueryCriteria(r.GetDefaultCriteria(), params ...)
	if err != nil {
		return nil, err
	}
	query := c.Find(conditions)
	var sorting = r.GetDefaultSorting()
	for _, param := range params {
		switch v := param.(type) {
		case bson.D, BinaryOperator, QueryBuilder:
			// Ignoring those types because they are exclusively used for
			// building the conditions to bootstrap the `mgo.Query` object.
		case *Sort:
			// For sorting it has a special case: it will aggregate all sortings
			// for a late modification.
			if sorting == nil {
				sorting = v.Fields
			} else {
				sorting = append(sorting, v.Fields...)
			}
		case QueryModifier:
			// Queryable objects will apply some transformation to the query.
			var err error
			query, err = v.Apply(query)
			if err != nil {
				return nil, err
			}
		default:
			// This type is not supported.
			return nil, NewErrTypeNotSupported(v)
		}
	}

	if len(sorting) > 0 {
		// If any sorting defined, applies it to the query.
		query = query.Sort(sorting...)
	}
	return query, nil
}

func NewErrTypeNotSupported(v interface{}) error {
	return fmt.Errorf("data type not supported: %s", reflect.TypeOf(v).String())
}
