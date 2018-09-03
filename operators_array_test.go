package repository_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"."
)

var _ = Describe("Operators Array", func() {
	BeforeEach(func() {
		Expect(clearSession()).To(BeNil())
	})

	Context("Operator: $elemMatch", func() {
		It("should find object with an $elemMatch condition single query", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.ElemMatch(
				"score", repository.LT("", 10),
			))).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
			Expect(objs[0].Score).To(Equal([]int{1, 2, 4, 5, 9}))
		})

		It("should fail object with an $elemMatch condition single query pass field failed", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.ElemMatch(
				"score", repository.LT("fail", 10),
			))).To(BeNil())
			Expect(objs).To(HaveLen(0))
		})

		It("should find object with an $elemMatch condition", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.ElemMatch(
				"score", repository.GTE("", 10),
				repository.LT("", 30),
			))).To(BeNil())
			Expect(objs).To(HaveLen(2))
			Expect(objs[0].Name).To(Equal("Scarlett"))
			Expect(objs[0].Score).To(Equal([]int{10, 20, 40}))
			Expect(objs[1].Name).To(Equal("Duke"))
			Expect(objs[1].Score).To(Equal([]int{10, 11}))
		})

		It("should find object with an $elemMatch condition match in array equals", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.ElemMatch(
				"details",
				repository.EQ("city", "Jamaica"),
				repository.EQ("fruit", "Pineapple"),
			))).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Scarlett"))
		})

		It("should find object with an $elemMatch condition match in array not equals", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.ElemMatch(
				"details",
				repository.EQ("city", "Brazil"),
				repository.Not(repository.EQ("fruit", "Apple")),
			))).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Duke"))
		})

		It("should find object with an $elemMatch condition match in array equals and not less", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.ElemMatch(
				"details",
				repository.EQ("city", "Brazil"),
				repository.Not(repository.LT("ability", 20)),
			))).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Duke"))
		})

		It("should find object with an $elemMatch condition match in array equals and greater", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.ElemMatch(
				"details",
				repository.EQ("city", "Jamaica"),
				repository.GTE("ability", 50),
			), repository.ElemMatch("score", repository.LT("", 20)))).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
		})

		It("should fail object with an $elemMatch condition needs an object", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			err := repository.FindAll(r, &objs, repository.ElemMatch(
				"name", "Jack",
			))
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("data type not supported:"))
		})
	})
})
