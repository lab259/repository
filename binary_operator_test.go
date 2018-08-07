package repository_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"."
	"./queries"
)

var _ = Describe("BinaryOperator", func() {
	BeforeEach(func() {
		Expect(clearSession()).To(BeNil())
	})

	It("should find objects with an $and condition", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(r, &objs, queries.And(
			queries.EQ("name", "Snake Eyes"),
			queries.EQ("age", 33),
		))).To(BeNil())
		Expect(objs).To(HaveLen(1))
		Expect(objs[0].Name).To(Equal("Snake Eyes"))
	})

	It("should find objects with an $or condition", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(r, &objs, queries.Or(
			queries.EQ("name", "Snake Eyes"),
			queries.EQ("name", "Duke"),
		))).To(BeNil())
		Expect(objs).To(HaveLen(2))
		Expect(objs[0].Name).To(Equal("Snake Eyes"))
		Expect(objs[1].Name).To(Equal("Duke"))
	})

	It("should find objects with combined $and inside $or", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(r, &objs, queries.Or(
			queries.EQ("age", 33),
			queries.And(
				queries.LT("strength", 7),
				queries.GT("agility", 8),
			),
		))).To(BeNil())
		Expect(objs).To(HaveLen(2))
		Expect(objs[0].Name).To(Equal("Snake Eyes"))
		Expect(objs[1].Name).To(Equal("Scarlett"))
	})

	It("should find objects with a criteria", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(r, &objs, queries.And(
			queries.WithCriteria(
				queries.LT("strength", 7),
				queries.GT("agility", 8),
			),
		))).To(BeNil())
		Expect(objs).To(HaveLen(1))
		Expect(objs[0].Name).To(Equal("Scarlett"))
	})

	It("should find objects with a criteria with a erroneous condition", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		err := repository.FindAll(r, &objs, queries.And(
			queries.WithCriteria(
				&erroneousCondition{},
			),
		))
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(ContainSubstring("forced error"))
	})
})
