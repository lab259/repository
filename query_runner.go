package repository

import "gopkg.in/mgo.v2"

type mgoDBRunner func(db *mgo.Database) error

// QueryRunner provide a infrastructure for executing querys in different
// contexts.
//
// In case of managing a single database, a simple `QueryRunner` implementation
// can be easily grab a connection from the pool and call the handler providing
// the DB reference.
//
// Additionaly, it also can implement be linked to a multi-tenanti environment
// that implements a database for each account. In that scenario, the
// `QueryRunner` implementation would make sure to take select the right
// database before calling the handler.
type QueryRunner interface {
	RunWithDB(handler mgoDBRunner) error
}
