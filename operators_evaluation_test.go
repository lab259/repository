package repository_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"."
)

var _ = Describe("Operators Evaluation", func() {
	BeforeEach(func() {
		Expect(clearSession()).To(BeNil())
	})

	It("should find objects with a $regex condition", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(r, &objs, repository.And(
			repository.WithCriteria(
				repository.Regex("name","Eyes","i"),
			),
		))).To(BeNil())
		Expect(objs).To(HaveLen(1))
		Expect(objs[0].Name).To(Equal("Snake Eyes"))
	})

	It("should find not exists objects with a $regex", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(r, &objs, repository.Not(
			repository.Regex("name", "Duke", ""),
		))).To(BeNil())
		Expect(objs).To(HaveLen(2))
		Expect(objs[0].Name).To(Equal("Snake Eyes"))
		Expect(objs[1].Name).To(Equal("Scarlett"))
	})

})
