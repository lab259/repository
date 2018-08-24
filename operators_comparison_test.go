package repository_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"."
)

var _ = Describe("Operators Comparison", func() {
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

	Context("Operator: $eq", func() {
		It("should find objects with an $eq condition", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.And(
				repository.WithCriteria(
					repository.EQ("strength", 5),
				),
			))).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Scarlett"))
			Expect(objs[0].Age).To(Equal(22))
			Expect(objs[0].Agiliy).To(Equal(9))
		})

		It("should find empty objects with an $eq condition", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.And(
				repository.WithCriteria(
					repository.EQ("strength", "fail"),
				),
			))).To(BeNil())
			Expect(objs).To(HaveLen(0))
		})
	})

	Context("Operator: $gt", func() {
		It("should find objects with a $gt condition", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.And(
				repository.WithCriteria(
					repository.GT("strength", 5),
				),
			))).To(BeNil())
			Expect(objs).To(HaveLen(2))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
			Expect(objs[1].Name).To(Equal("Duke"))
		})

		It("should find empty objects with a $gt condition", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.And(
				repository.WithCriteria(
					repository.GT("strength", nil),
				),
			))).To(BeNil())
			Expect(objs).To(HaveLen(0))
		})
	})

	Context("Operator: $gte", func() {
		It("should find objects with a $gte condition", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.And(
				repository.WithCriteria(
					repository.GTE("strength", 5),
				),
			))).To(BeNil())
			Expect(objs).To(HaveLen(3))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
			Expect(objs[1].Name).To(Equal("Scarlett"))
			Expect(objs[2].Name).To(Equal("Duke"))
		})

		It("should find empty objects with a $gte condition", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.And(
				repository.WithCriteria(
					repository.GTE("age", 5000),
				),
			))).To(BeNil())
			Expect(objs).To(HaveLen(0))
		})
	})

	Context("Operator: $in", func() {
		It("should find objects with an $in condition", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.And(
				repository.WithCriteria(
					repository.IN("tags", []string{"red", "black"}),
				),
			))).To(BeNil())
			Expect(objs).To(HaveLen(2))
			Expect(objs[0].Name).To(Equal("Scarlett"))
			Expect(objs[0].Tags).To(Equal([]string{"yellow", "red"}))
			Expect(objs[1].Name).To(Equal("Duke"))
			Expect(objs[1].Tags).To(Equal([]string{"green", "black"}))
		})

		It("should find empty objects with an $in condition", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.And(
				repository.WithCriteria(
					repository.IN("tags", []string{"gray"}),
				),
			))).To(BeNil())
			Expect(objs).To(HaveLen(0))
		})
	})

	Context("Operator: $lt", func() {
		It("should find objects with a $lt condition", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.And(
				repository.WithCriteria(
					repository.LT("age", 33),
				),
			))).To(BeNil())
			Expect(objs).To(HaveLen(2))
			Expect(objs[0].Name).To(Equal("Scarlett"))
			Expect(objs[1].Name).To(Equal("Duke"))
		})

		It("should find empty objects with a $lt condition", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.And(
				repository.WithCriteria(
					repository.LT("age", 21),
				),
			))).To(BeNil())
			Expect(objs).To(HaveLen(0))
		})
	})

	Context("Operator: $lt", func() {
		It("should find objects with a $lte condition", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.And(
				repository.WithCriteria(
					repository.LTE("age", 33),
				),
			))).To(BeNil())
			Expect(objs).To(HaveLen(3))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
			Expect(objs[1].Name).To(Equal("Scarlett"))
			Expect(objs[2].Name).To(Equal("Duke"))
		})

		It("should find empty objects with a $gte condition", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.And(
				repository.WithCriteria(
					repository.LTE("age", 21),
				),
			))).To(BeNil())
			Expect(objs).To(HaveLen(0))
		})
	})

	Context("Operator: $ne", func() {
		It("should find objects with an $ne condition", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.And(
				repository.WithCriteria(
					repository.NE("name", "Scarlett"),
				),
			))).To(BeNil())
			Expect(objs).To(HaveLen(2))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
			Expect(objs[1].Name).To(Equal("Duke"))
		})

		It("should find objects with an $ne condition and not value match", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.And(
				repository.WithCriteria(
					repository.NE("name", "Jack"),
				),
			))).To(BeNil())
			Expect(objs).To(HaveLen(3))
		})
	})

	Context("Operator: $nin", func() {
		It("should find objects with a $nin condition", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.And(
				repository.WithCriteria(
					repository.NIN("tags", []string{"red", "black"}),
				),
			))).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
			Expect(objs[0].Tags).To(Equal([]string{"blue", "yellow", "green"}))
		})

		It("should find empty objects with a $nin condition", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			err := repository.FindAll(r, &objs, repository.And(
				repository.WithCriteria(
					repository.NIN("tags", "gray"),
				),
			))
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("$nin needs an array"))
		})

	})
})
