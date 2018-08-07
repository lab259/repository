package repository_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/mgo.v2/bson"

	"."
)

var _ = Describe("FindAll", func() {
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
		Expect(repository.FindAll(r, &objs, repository.WithCriteria(repository.LT("age", 30)))).To(BeNil())
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
		Expect(repository.FindAll(r, &objs, repository.WithCriteria(repository.LT("age", 30)), repository.WithSort("name"))).To(BeNil())
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
		Expect(repository.FindAll(r, &objs, repository.WithPage(1, 2))).To(BeNil())
		Expect(objs).To(HaveLen(1))
		Expect(objs[0].Name).To(Equal("Duke"))
		Expect(objs[0].Age).To(Equal(22))
	})

	It("should find all objects with skip", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(r, &objs, repository.Skip(1))).To(BeNil())
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
		Expect(repository.FindAll(r, &objs, repository.Limit(1), repository.WithSort("name"))).To(BeNil())
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
			repository.Limit(1),
			repository.WithSort("name"),
			repository.LT("age", 30),
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
			repository.Limit(2),
			repository.WithSort("name"),
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
				repository.WithCriteria(
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
				repository.WithCriteria(
					repository.And(
						repository.LT("age", 35),
						repository.GT("age", 30),
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
				repository.WithCriteria(
					repository.And(
						repository.LT("age", 35),
						repository.GT("age", 30),
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
				repository.WithCriteria(bson.M{
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
				repository.WithCriteria(
					&erroneousCondition{},
				),
			)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("forced error"))
		})
	})
})
