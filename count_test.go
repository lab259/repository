package repository_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"."
)

var _ = Describe("Count", func() {
	BeforeEach(func() {
		Expect(clearSession()).To(BeNil())
	})

	It("should count all objects", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		count, err := repository.Count(r)
		Expect(err).To(BeNil())
		Expect(count).To(Equal(3))
	})

	It("should count objects with some criteria", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		count, err := repository.Count(r, repository.WithCriteria(repository.LT("age", 30)))
		Expect(err).To(BeNil())
		Expect(count).To(Equal(2))
	})

	It("should find all objects with sorting", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		count, err := repository.Count(r, repository.WithCriteria(repository.LT("age", 30)), repository.WithSort("name"))
		Expect(err).To(BeNil())
		Expect(count).To(Equal(2))
	})
})
