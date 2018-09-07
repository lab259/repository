package repository_test

import (
	"log"
	"testing"

	"."
	"github.com/globalsign/mgo"
	"github.com/jamillosantos/macchiato"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

type mgoService struct {
}

func (service *mgoService) RunWithSession(handler func(session *mgo.Session) error) error {
	sess := session.Clone()
	defer sess.Close()
	return handler(session)
}

var (
	session            *mgo.Session
	defaultQueryRunner repository.QueryRunner
)

func connect() {
	s, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{"localhost"},
		Database: "repository_test",
	})

	gomega.Expect(err).To(gomega.BeNil())
	session = s
	defaultQueryRunner = repository.NewSimpleQueryRunner(&mgoService{}, "repository_test")
}

func TestRepository(t *testing.T) {
	log.SetOutput(ginkgo.GinkgoWriter)
	gomega.RegisterFailHandler(ginkgo.Fail)
	connect()
	macchiato.RunSpecs(t, "Repository Test Suite")
}

func clearSession() error {
	return RunWithSession(func(db *mgo.Database) error {
		collections, err := db.CollectionNames()
		gomega.Expect(err).To(gomega.BeNil())
		for _, c := range collections {
			gomega.Expect(db.C(c).DropCollection()).To(gomega.BeNil())
		}
		return nil
	})
}

func RunWithSession(handler func(db *mgo.Database) error) error {
	sess := session.Clone()
	defer sess.Close()
	return handler(session.DB(""))
}
