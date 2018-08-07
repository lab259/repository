package repository_test

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/jamillosantos/macchiato"
	"testing"
	"log"
	"gopkg.in/mgo.v2"
	"."
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

/*
type testRepNoDefaultCriteriaNoDefaultSorting struct {
}

func (rep *testRepNoDefaultCriteriaNoDefaultSorting) GetCollectionName() string {
	return "collection"
}

func (rep *testRepNoDefaultCriteriaNoDefaultSorting) GetQueryRunner() repository.QueryRunner {
	return defaultQueryRunner
}

func (rep *testRepNoDefaultCriteriaNoDefaultSorting) GetDefaultCriteria() bson.D {
	return nil
}

func (rep *testRepNoDefaultCriteriaNoDefaultSorting) GetDefaultSorting() []string {
	return nil
}
*/

func TestRepository(t *testing.T) {
	log.SetOutput(ginkgo.GinkgoWriter)
	gomega.RegisterFailHandler(ginkgo.Fail)
	s, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{"localhost"},
		Database: "repository_test",
	})
	if err != nil {
		t.Fatalf("error initializing the session: %s", err)
	}
	session = s
	defaultQueryRunner = repository.NewSimpleQueryRunner(&mgoService{}, "repository_test")

	/*
	type testRepObject struct {
		ID   bson.ObjectId `bson:"_id,omitempty"`
		Name string        `bson:"name,omitempty"`
		Age  int           `bson:"age,omitempty"`
	}
	r := &testRepNoDefaultCriteriaNoDefaultSorting{}
	tobj := &testRepObject{
		ID:   bson.NewObjectId(),
		Name: "Snake Eyes",
		Age:  33,
	}
	err = repository.Create(r, tobj)
	if err != nil {
		t.FailNow()
	}
	var obj testRepObject
	repository.Find(r, &obj, queries.WithCriteria(
		queries.And(
			queries.EQ("name", "Snake Eyes 1"),
			queries.GT("age", 30),
		),
	))
	*/
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
