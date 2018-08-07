package repository_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
		"."
	"./queries"
)

var _ = Describe("CountAndFind", func() {
	BeforeEach(func() {
		Expect(clearSession()).To(BeNil())
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
		count, err := repository.CountAndFindAll(r, &objs, queries.WithCriteria(queries.LT("age", 30)))
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
		count, err := repository.CountAndFindAll(r, &objs, queries.WithCriteria(queries.LT("age", 30)), queries.Sort("name"))
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
		count, err := repository.CountAndFindAll(r, &objs, queries.WithPage(1, 2))
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
		count, err := repository.CountAndFindAll(r, &objs, queries.Skip(1))
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
		count, err := repository.CountAndFindAll(r, &objs, queries.Limit(1), queries.Sort("name"))
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
