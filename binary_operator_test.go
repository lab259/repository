package repository_test

import (
	"github.com/lab259/repository"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BinaryOperator", func() {
	BeforeEach(func() {
		Expect(clearSession()).To(BeNil())
	})

	It("should find objects with an $and condition", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(r, &objs, repository.And(
			repository.EQ("name", "Snake Eyes"),
			repository.EQ("age", 33),
		))).To(BeNil())
		Expect(objs).To(HaveLen(1))
		Expect(objs[0].Name).To(Equal("Snake Eyes"))
	})

	It("should find objects with an $or condition", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(r, &objs, repository.Or(
			repository.EQ("name", "Snake Eyes"),
			repository.EQ("name", "Duke"),
		))).To(BeNil())
		Expect(objs).To(HaveLen(2))
		Expect(objs[0].Name).To(Equal("Snake Eyes"))
		Expect(objs[1].Name).To(Equal("Duke"))
	})

	It("should find objects with combined $and inside $or", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(r, &objs, repository.Or(
			repository.EQ("age", 33),
			repository.And(
				repository.LT("strength", 7),
				repository.GT("agility", 8),
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
		Expect(repository.FindAll(r, &objs, repository.And(
			repository.WithCriteria(
				repository.LT("strength", 7),
				repository.GT("agility", 8),
			),
		))).To(BeNil())
		Expect(objs).To(HaveLen(1))
		Expect(objs[0].Name).To(Equal("Scarlett"))
	})

	It("should find objects with a criteria with a erroneous condition", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		err := repository.FindAll(r, &objs, repository.And(
			repository.WithCriteria(
				&erroneousCondition{},
			),
		))
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(ContainSubstring("forced error"))
	})

	It("should find objects with an $not condition", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(r, &objs, repository.Not(
			repository.EQ("name", "Duke"),
		),
		)).To(BeNil())
		Expect(objs).To(HaveLen(2))
		Expect(objs[0].Name).To(Equal("Snake Eyes"))
		Expect(objs[1].Name).To(Equal("Scarlett"))
	})

	It("should find objects with an $nor condition", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(r, &objs, repository.Nor("name",
			repository.WithCriteria(
				repository.EQ("name", "Snake Eyes"),
				repository.NE("strength", 8),
			),
		))).To(BeNil())
		Expect(objs).To(HaveLen(2))
		Expect(objs[0].Name).To(Equal("Scarlett"))
	})
})
