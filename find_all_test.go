package repository_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/mgo.v2/bson"

	"."
	"./queries"
	"errors"
	"gopkg.in/mgo.v2"
)

type erroneousCondition struct{}

func (err *erroneousCondition) GetCondition() (bson.DocElem, error) {
	return bson.DocElem{}, errors.New("forced error")
}

type erroneousQueryModifier struct{}

func (err *erroneousQueryModifier) Apply(query *mgo.Query) (*mgo.Query, error) {
	return nil, errors.New("forced error")
}

var _ = Describe("FindAll", func() {
	insertObjects := func(r repository.Repository) (bson.ObjectId, bson.ObjectId, bson.ObjectId) {
		objid1, objid2, objid3 := bson.NewObjectId(), bson.NewObjectId(), bson.NewObjectId()
		Expect(repository.Create(r, &testRepObject{
			ID:   objid1,
			Name: "Snake Eyes",
			Age:  33,
		})).To(BeNil())
		Expect(repository.Create(r, &testRepObject{
			ID:   objid2,
			Name: "Scarlett",
			Age:  22,
		})).To(BeNil())
		Expect(repository.Create(r, &testRepObject{
			ID:   objid3,
			Name: "Duke",
			Age:  22,
		})).To(BeNil())

		return objid1, objid2, objid3
	}

	BeforeEach(func() {
		Expect(clearSession()).To(BeNil())
	})

	It("should find all objects", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(r, &objs)).To(BeNil())
		Expect(objs).To(HaveLen(3))
		Expect(objs[0].Name).To(Equal("Snake Eyes"))
		Expect(objs[0].Age).To(Equal(33))
		Expect(objs[1].Name).To(Equal("Scarlett"))
		Expect(objs[1].Age).To(Equal(22))
		Expect(objs[2].Name).To(Equal("Duke"))
		Expect(objs[2].Age).To(Equal(22))
	})

	It("should find all objects with some criteria", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(r, &objs, queries.WithCriteria(queries.LT("age", 30)))).To(BeNil())
		Expect(objs).To(HaveLen(2))
		Expect(objs[0].Name).To(Equal("Scarlett"))
		Expect(objs[0].Age).To(Equal(22))
		Expect(objs[1].Name).To(Equal("Duke"))
		Expect(objs[1].Age).To(Equal(22))
	})

	It("should find all objects with sorting", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(r, &objs, queries.WithCriteria(queries.LT("age", 30)), queries.Sort("name"))).To(BeNil())
		Expect(objs).To(HaveLen(2))
		Expect(objs[0].Name).To(Equal("Duke"))
		Expect(objs[0].Age).To(Equal(22))
		Expect(objs[1].Name).To(Equal("Scarlett"))
		Expect(objs[1].Age).To(Equal(22))
	})

	It("should find all objects with paging", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(r, &objs, queries.WithPage(1, 2))).To(BeNil())
		Expect(objs).To(HaveLen(1))
		Expect(objs[0].Name).To(Equal("Duke"))
		Expect(objs[0].Age).To(Equal(22))
	})

	It("should find all objects with skip", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(r, &objs, queries.Skip(1))).To(BeNil())
		Expect(objs).To(HaveLen(2))
		Expect(objs[0].Name).To(Equal("Scarlett"))
		Expect(objs[0].Age).To(Equal(22))
		Expect(objs[1].Name).To(Equal("Duke"))
		Expect(objs[1].Age).To(Equal(22))
	})

	It("should find all objects with limit", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(r, &objs, queries.Limit(1), queries.Sort("name"))).To(BeNil())
		Expect(objs).To(HaveLen(1))
		Expect(objs[0].Name).To(Equal("Duke"))
		Expect(objs[0].Age).To(Equal(22))
	})

	It("should find all with criteria, sort and limit", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(
			r,
			&objs,
			queries.Limit(1),
			queries.Sort("name"),
			queries.LT("age", 30),
		)).To(BeNil())
		Expect(objs).To(HaveLen(1))
		Expect(objs[0].Name).To(Equal("Duke"))
		Expect(objs[0].Age).To(Equal(22))
	})

	It("should find all with default criteria", func() {
		r := &testRepWithDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(
			r,
			&objs,
		)).To(BeNil())
		Expect(objs).To(HaveLen(2))
		Expect(objs[0].Name).To(Equal("Scarlett"))
		Expect(objs[0].Age).To(Equal(22))
		Expect(objs[1].Name).To(Equal("Duke"))
		Expect(objs[1].Age).To(Equal(22))
	})

	It("should find all with default criteria", func() {
		r := &testRepWithDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(
			r,
			&objs,
		)).To(BeNil())
		Expect(objs).To(HaveLen(2))
		Expect(objs[0].Name).To(Equal("Scarlett"))
		Expect(objs[0].Age).To(Equal(22))
		Expect(objs[1].Name).To(Equal("Duke"))
		Expect(objs[1].Age).To(Equal(22))
	})

	It("should find all with default criteria", func() {
		r := &testRepWithDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(
			r,
			&objs,
		)).To(BeNil())
		Expect(objs).To(HaveLen(2))
		Expect(objs[0].Name).To(Equal("Scarlett"))
		Expect(objs[0].Age).To(Equal(22))
		Expect(objs[1].Name).To(Equal("Duke"))
		Expect(objs[1].Age).To(Equal(22))
	})

	It("should find all with default criteria and additional sorting", func() {
		r := &testRepNoDefaultCriteriaWithDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(
			r,
			&objs,
			queries.Limit(2),
			queries.Sort("name"),
		)).To(BeNil())
		Expect(objs).To(HaveLen(2))
		Expect(objs[0].Name).To(Equal("Duke"))
		Expect(objs[0].Age).To(Equal(22))
		Expect(objs[1].Name).To(Equal("Scarlett"))
		Expect(objs[1].Age).To(Equal(22))
	})

	It("should find objects with a RawQuery", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		objid1, objid2, objid3 := insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(
			r,
			&objs,
			&repository.RawCriteria{
				bson.DocElem{Name: "_id", Value: bson.D{bson.DocElem{Name: "$in", Value: []interface{}{objid1, objid2, objid3}}}},
			},
		)).To(BeNil())
		Expect(objs).To(HaveLen(3))
		Expect(objs[0].Name).To(Equal("Snake Eyes"))
		Expect(objs[1].Name).To(Equal("Scarlett"))
		Expect(objs[2].Name).To(Equal("Duke"))
	})

	It("should find objects with a bson.D", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		objid1, _, _ := insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(
			r,
			&objs,
			bson.D{
				bson.DocElem{Name: "_id", Value: bson.D{bson.DocElem{Name: "$in", Value: []interface{}{objid1}}}},
			},
		)).To(BeNil())
		Expect(objs).To(HaveLen(1))
		Expect(objs[0].Name).To(Equal("Snake Eyes"))
	})

	Describe("WithCriteria", func() {
		It("should find objects with a Criteria with a bson.DocElem inside", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			objid1, _, _ := insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(
				r,
				&objs,
				queries.WithCriteria(
					bson.DocElem{Name: "_id", Value: bson.D{bson.DocElem{Name: "$in", Value: []interface{}{objid1}}}},
				),
			)).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
		})

		It("should find objects with a Criteria with a BooleanOperator inside", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(
				r,
				&objs,
				queries.WithCriteria(
					queries.And(
						queries.LT("age", 35),
						queries.GT("age", 30),
					),
				),
			)).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
		})

		It("should find objects with a Criteria with a BooleanOperator inside", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(
				r,
				&objs,
				queries.WithCriteria(
					queries.And(
						queries.LT("age", 35),
						queries.GT("age", 30),
					),
				),
			)).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
		})

		It("should fail with a criteria that is not supported", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			err := repository.FindAll(
				r,
				&objs,
				queries.WithCriteria(bson.M{
					"name": "Snake Eyes",
				}),
			)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("data type not supported"))
		})

		It("should fail with a criteria with a condition that returns error", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			err := repository.FindAll(
				r,
				&objs,
				queries.WithCriteria(
					&erroneousCondition{},
				),
			)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("forced error"))
		})
	})
})
