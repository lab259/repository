package repository_test

import (
	"errors"

	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"
	"github.com/lab259/repository"

	. "github.com/onsi/gomega"
)

type testRepObject struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `bson:"name,omitempty"`
	Age      int           `bson:"age,omitempty"`
	Strength int           `bson:"strength,omitempty"`
	Agility  int           `bson:"agility,omitempty"`
	Tags     []string      `bson:"tags,omitempty"`
	Status   bool          `bson:"status,omitempty"`
	Score    []int         `bson:"score,omitempty"`
	Details  []bson.M      `bson:"details,omitempty"`
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

func createIndexes(r repository.Repository) {
	index := mgo.Index{
		Key: []string{"$text:name"},
	}
	Expect(r.GetQueryRunner().RunWithDB(func(db *mgo.Database) error {
		return db.C(r.GetCollectionName()).EnsureIndex(index)
	})).To(BeNil())
}

func insertObjects(r repository.Repository) (bson.ObjectId, bson.ObjectId, bson.ObjectId) {
	objid1, objid2, objid3 := bson.NewObjectId(), bson.NewObjectId(), bson.NewObjectId()

	Expect(repository.Create(r, &testRepObject{
		ID:       objid1,
		Name:     "Snake Eyes",
		Age:      33,
		Strength: 7,
		Agility:  9,
		Tags:     []string{"blue", "yellow", "green"},
		Status:   true,
		Score:    []int{1, 2, 4, 5, 9},
		Details: []bson.M{
			{
				"ability": 50,
				"fruit":   "Orange",
				"city":    "Jamaica",
			},
			{
				"ability": 10,
				"fruit":   "Apple",
				"city":    "Brazil",
			},
			{
				"ability": 30,
				"fruit":   "Lime",
				"city":    "News York",
			},
		},
	})).To(BeNil())
	Expect(repository.Create(r, &testRepObject{
		ID:       objid2,
		Name:     "Scarlett",
		Age:      22,
		Strength: 5,
		Agility:  9,
		Tags:     []string{"yellow", "red"},
		Status:   false,
		Score:    []int{10, 20, 40},
		Details: []bson.M{
			{
				"ability": 30,
				"fruit":   "Pineapple",
				"city":    "Jamaica",
			},
			{
				"ability": 40,
				"fruit":   "Apple",
				"city":    "Canada",
			},
		},
	})).To(BeNil())
	Expect(repository.Create(r, &testRepObject{
		ID:       objid3,
		Name:     "Duke",
		Age:      22,
		Strength: 8,
		Agility:  7,
		Tags:     []string{"green", "black"},
		Score:    []int{10, 11},
		Details: []bson.M{
			{
				"ability": 20,
				"fruit":   "Guava",
				"city":    "Brazil",
			},
		},
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
