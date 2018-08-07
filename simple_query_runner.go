package repository

import (
	"gopkg.in/mgo.v2"
)

type MgoServiceRunner interface {
	RunWithSession(handler func(session *mgo.Session) error) error
}

type SimpleQueryRunner struct {
	service  MgoServiceRunner
	database string
}

func NewSimpleQueryRunner(service MgoServiceRunner, database string) *SimpleQueryRunner {
	return &SimpleQueryRunner{
		service:  service,
		database: database,
	}
}

// RunWithDB runs
func (runner *SimpleQueryRunner) RunWithDB(handler func (db *mgo.Database) error) error {
	return runner.service.RunWithSession(func(session *mgo.Session) error {
		return handler(session.DB(runner.database))
	})
}
