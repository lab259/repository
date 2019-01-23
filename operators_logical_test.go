package repository_test

import (
	"github.com/lab259/repository"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Operators Logical", func() {
	BeforeEach(func() {
		Expect(clearSession()).To(BeNil())
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

	Context("Operator: $and", func() {
		It("should find objects with an $and condition", func() {
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
	})

	Context("Operator: $nor", func() {
		It("should find objects with an $nor condition", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.Nor(
				repository.WithCriteria(
					repository.EQ("strength", 5),
				),
			))).To(BeNil())
			Expect(objs).To(HaveLen(2))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
		})

		It("should find objects with an $nor and a criteria with a erroneous condition", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			err := repository.FindAll(r, &objs, repository.Nor(
				repository.WithCriteria(
					&erroneousCondition{},
				),
			))
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("forced error"))
		})
	})

	Context("Operator: $or", func() {
		It("should find objects with an $or condition", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.Or(
				repository.WithCriteria(
					repository.NE("strength", 5),
					repository.EQ("age", 33),
				),
			))).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
			Expect(objs[0].Age).To(Equal(33))
			Expect(objs[0].Strength).To(Equal(7))
			Expect(objs[0].Agility).To(Equal(9))
		})

		It("should find objects with an $or and a criteria with a erroneous condition", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			err := repository.FindAll(r, &objs, repository.Or(
				repository.WithCriteria(
					&erroneousCondition{},
				),
			))
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("forced error"))
		})
	})

	Context("Operator: $not", func() {
		It("should find objects with an $not condition", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.Not(
				repository.LT("strength", 7),
			))).To(BeNil())
			Expect(objs).To(HaveLen(2))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
			Expect(objs[0].Age).To(Equal(33))
			Expect(objs[0].Strength).To(Equal(7))
			Expect(objs[0].Agility).To(Equal(9))
		})

		It("should find objects with an $not and a criteria with a erroneous condition", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			err := repository.FindAll(r, &objs, repository.Not(
				repository.GT("", ""),
			))
			Expect(err)
		})
	})

})
