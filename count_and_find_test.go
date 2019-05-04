package repository_test

import (
	"github.com/lab259/repository"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CountAndFind", func() {
	BeforeEach(func() {
		Expect(clearSession()).To(BeNil())
	})

	It("should find all objects (default repository)", func() {
		r := NewRepository()
		insertObjects(r)
		objs := make([]testRepObject, 0)
		count, err := r.CountAndFindAll(&objs)
		Expect(err).To(BeNil())
		Expect(count).To(Equal(3))
		Expect(objs).To(HaveLen(3))
		Expect(objs[0].Name).To(Equal("Snake Eyes"))
		Expect(objs[0].Age).To(Equal(33))
		Expect(objs[1].Name).To(Equal("Scarlett"))
		Expect(objs[1].Age).To(Equal(22))
		Expect(objs[2].Name).To(Equal("Duke"))
		Expect(objs[2].Age).To(Equal(22))
	})

	It("should find all objects", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		count, err := repository.CountAndFindAll(r, &objs)
		Expect(err).To(BeNil())
		Expect(count).To(Equal(3))
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
		count, err := repository.CountAndFindAll(r, &objs, repository.WithCriteria(repository.LT("age", 30)))
		Expect(err).To(BeNil())
		Expect(count).To(Equal(2))
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
		count, err := repository.CountAndFindAll(r, &objs, repository.WithCriteria(repository.LT("age", 30)), repository.WithSort("name"))
		Expect(err).To(BeNil())
		Expect(count).To(Equal(2))
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
		count, err := repository.CountAndFindAll(r, &objs, repository.WithPage(1, 2))
		Expect(err).To(BeNil())
		Expect(count).To(Equal(3))
		Expect(objs).To(HaveLen(1))
		Expect(objs[0].Name).To(Equal("Duke"))
		Expect(objs[0].Age).To(Equal(22))
	})

	It("should find all objects with skip", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		count, err := repository.CountAndFindAll(r, &objs, repository.Skip(1))
		Expect(err).To(BeNil())
		Expect(count).To(Equal(3))
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
		count, err := repository.CountAndFindAll(r, &objs, repository.Limit(1), repository.WithSort("name"))
		Expect(err).To(BeNil())
		Expect(count).To(Equal(3))
		Expect(objs).To(HaveLen(1))
		Expect(objs[0].Name).To(Equal("Duke"))
		Expect(objs[0].Age).To(Equal(22))
	})

	It("should find all with default criteria", func() {
		r := &testRepWithDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]testRepObject, 0)
		count, err := repository.CountAndFindAll(
			r,
			&objs,
		)
		Expect(err).To(BeNil())
		Expect(count).To(Equal(2))
		Expect(objs).To(HaveLen(2))
		Expect(objs[0].Name).To(Equal("Scarlett"))
		Expect(objs[0].Age).To(Equal(22))
		Expect(objs[1].Name).To(Equal("Duke"))
		Expect(objs[1].Age).To(Equal(22))
	})
})
