package repository_test

import (
	"github.com/globalsign/mgo/bson"
	. "github.com/onsi/gomega"

	"errors"

	"."
	"github.com/globalsign/mgo"
)

type testRepObject struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `bson:"name,omitempty"`
	Age      int           `bson:"age,omitempty"`
	Strength int           `bson:"strength,omitempty"`
	Agiliy   int           `bson:"agility,omitempty"`
}

type testRepNoDefaultCriteriaNoDefaultSorting struct {
}

func (rep *testRepNoDefaultCriteriaNoDefaultSorting) GetCollectionName() string {
	return "collection"
}

func (rep *testRepNoDefaultCriteriaNoDefaultSorting) GetQueryRunner() repository.QueryRunner {
	return defaultQueryRunner
}

func (rep *testRepNoDefaultCriteriaNoDefaultSorting) GetDefaultCriteria() interface{} {
	return nil
}

func (rep *testRepNoDefaultCriteriaNoDefaultSorting) GetDefaultSorting() []string {
	return nil
}

type testRepWithDefaultCriteriaNoDefaultSorting struct {
}

func (rep *testRepWithDefaultCriteriaNoDefaultSorting) GetCollectionName() string {
	return "collection"
}

func (rep *testRepWithDefaultCriteriaNoDefaultSorting) GetQueryRunner() repository.QueryRunner {
	return defaultQueryRunner
}

func (rep *testRepWithDefaultCriteriaNoDefaultSorting) GetDefaultCriteria() interface{} {
	return repository.LT("age", 30)
}

func (rep *testRepWithDefaultCriteriaNoDefaultSorting) GetDefaultSorting() []string {
	return nil
}

type testRepNoDefaultCriteriaWithDefaultSorting struct {
}

func (rep *testRepNoDefaultCriteriaWithDefaultSorting) GetCollectionName() string {
	return "collection"
}

func (rep *testRepNoDefaultCriteriaWithDefaultSorting) GetQueryRunner() repository.QueryRunner {
	return defaultQueryRunner
}

func (rep *testRepNoDefaultCriteriaWithDefaultSorting) GetDefaultCriteria() interface{} {
	return nil
}

func (rep *testRepNoDefaultCriteriaWithDefaultSorting) GetDefaultSorting() []string {
	return []string{"age"}
}

func insertObjects(r repository.Repository) (bson.ObjectId, bson.ObjectId, bson.ObjectId) {
	objid1, objid2, objid3 := bson.NewObjectId(), bson.NewObjectId(), bson.NewObjectId()
	Expect(repository.Create(r, &testRepObject{
		ID:       objid1,
		Name:     "Snake Eyes",
		Age:      33,
		Strength: 7,
		Agiliy:   9,
	})).To(BeNil())
	Expect(repository.Create(r, &testRepObject{
		ID:       objid2,
		Name:     "Scarlett",
		Age:      22,
		Strength: 5,
		Agiliy:   9,
	})).To(BeNil())
	Expect(repository.Create(r, &testRepObject{
		ID:       objid3,
		Name:     "Duke",
		Age:      22,
		Strength: 8,
		Agiliy:   7,
	})).To(BeNil())

	return objid1, objid2, objid3
}

type erroneousCondition struct{}

func (err *erroneousCondition) GetCondition() (bson.DocElem, error) {
	return bson.DocElem{}, errors.New("forced error")
}

type erroneousQueryModifier struct{}

func (err *erroneousQueryModifier) Apply(query *mgo.Query) (*mgo.Query, error) {
	return nil, errors.New("forced error")
}
