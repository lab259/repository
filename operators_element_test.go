package repository_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"."
)

var _ = Describe("Operators Element", func() {
	BeforeEach(func() {
		Expect(clearSession()).To(BeNil())
	})

	It("should find objects with an $exists condition", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(r, &objs, repository.And(
			repository.WithCriteria(
				repository.Exists("status", true),
			),
		))).To(BeNil())
		Expect(objs).To(HaveLen(1))
		Expect(objs[0].Name).To(Equal("Snake Eyes"))
	})

	It("should find not exists objects with an $exists", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(r, &objs, repository.Not(
			repository.Exists("status", true),
		))).To(BeNil())
		Expect(objs).To(HaveLen(2))
		Expect(objs[0].Name).To(Equal("Scarlett"))
		Expect(objs[1].Name).To(Equal("Duke"))
	})

})
