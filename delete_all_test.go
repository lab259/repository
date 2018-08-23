package repository_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
		"."
	)

var _ = Describe("DeleteAll", func() {
	BeforeEach(func() {
		Expect(clearSession()).To(BeNil())
	})

	It("should delete an object", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		_, objid2, _ := insertObjects(r)
		deleted, err := repository.DeleteAll(r, repository.ByID(objid2))
		Expect(err).To(BeNil())
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(r, &objs)).To(BeNil())
		Expect(deleted).To(Equal(1))
		Expect(objs).To(HaveLen(2))
		Expect(objs[0].Name).To(Equal("Snake Eyes"))
		Expect(objs[1].Name).To(Equal("Duke"))
	})

	It("should delete base on a criteria", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		deleted, err := repository.DeleteAll(r, repository.LT("age", 30))
		Expect(err).To(BeNil())
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(r, &objs)).To(BeNil())
		Expect(deleted).To(Equal(2))
		Expect(objs).To(HaveLen(1))
		Expect(objs[0].Name).To(Equal("Snake Eyes"))
	})
})
