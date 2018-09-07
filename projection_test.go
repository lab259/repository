package repository_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"."
	"gopkg.in/mgo.v2/bson"
)

var _ = Describe("Projection", func() {
	BeforeEach(func() {
		Expect(clearSession()).To(BeNil())
	})

	It("should project fields using strings with no string modifier", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]bson.M, 0)
		err := repository.FindAll(
			r,
			&objs,
			repository.Projection("name"),
		)
		Expect(err).To(BeNil())
		Expect(objs).To(HaveLen(3))
		Expect(objs[0]).To(HaveLen(2))
		Expect(objs[0]).To(HaveKey("_id"))
		Expect(objs[0]).To(HaveKeyWithValue("name", "Snake Eyes"))
		Expect(objs[1]).To(HaveLen(2))
		Expect(objs[1]).To(HaveKey("_id"))
		Expect(objs[1]).To(HaveKeyWithValue("name", "Scarlett"))
		Expect(objs[2]).To(HaveLen(2))
		Expect(objs[2]).To(HaveKey("_id"))
		Expect(objs[2]).To(HaveKeyWithValue("name", "Duke"))
	})

	It("should project fields using strings with a plus modifier", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]bson.M, 0)
		err := repository.FindAll(
			r,
			&objs,
			repository.Projection("+name", "+age"),
		)
		Expect(err).To(BeNil())
		Expect(objs).To(HaveLen(3))
		Expect(objs[0]).To(HaveLen(3))
		Expect(objs[0]).To(HaveKey("_id"))
		Expect(objs[0]).To(HaveKeyWithValue("name", "Snake Eyes"))
		Expect(objs[0]).To(HaveKeyWithValue("age", 33))
		Expect(objs[1]).To(HaveLen(3))
		Expect(objs[1]).To(HaveKey("_id"))
		Expect(objs[1]).To(HaveKeyWithValue("name", "Scarlett"))
		Expect(objs[1]).To(HaveKeyWithValue("age", 22))
		Expect(objs[2]).To(HaveLen(3))
		Expect(objs[2]).To(HaveKey("_id"))
		Expect(objs[2]).To(HaveKeyWithValue("name", "Duke"))
		Expect(objs[2]).To(HaveKeyWithValue("age", 22))
	})

	It("should project fields using strings with a minus modifier", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]bson.M, 0)
		err := repository.FindAll(
			r,
			&objs,
			repository.Projection("-_id", "-age", "-score", "-status", "-details", "-strength", "-agility", "-tags"),
		)
		Expect(err).To(BeNil())
		Expect(objs).To(HaveLen(3))
		Expect(objs[0]).To(HaveLen(1))
		Expect(objs[0]).To(HaveKeyWithValue("name", "Snake Eyes"))
		Expect(objs[1]).To(HaveLen(1))
		Expect(objs[1]).To(HaveKeyWithValue("name", "Scarlett"))
		Expect(objs[2]).To(HaveLen(1))
		Expect(objs[2]).To(HaveKeyWithValue("name", "Duke"))
	})

	It("should project fields using DocElem", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]bson.M, 0)
		err := repository.FindAll(
			r,
			&objs,
			repository.Projection(bson.DocElem{Name: "name", Value: 1}),
		)
		Expect(err).To(BeNil())
		Expect(objs).To(HaveLen(3))
		Expect(objs[0]).To(HaveLen(2))
		Expect(objs[0]).To(HaveKey("_id"))
		Expect(objs[0]).To(HaveKeyWithValue("name", "Snake Eyes"))
		Expect(objs[1]).To(HaveLen(2))
		Expect(objs[1]).To(HaveKey("_id"))
		Expect(objs[1]).To(HaveKeyWithValue("name", "Scarlett"))
		Expect(objs[2]).To(HaveLen(2))
		Expect(objs[2]).To(HaveKey("_id"))
		Expect(objs[2]).To(HaveKeyWithValue("name", "Duke"))
	})

	It("should project fields using *DocElem", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		objs := make([]bson.M, 0)
		err := repository.FindAll(
			r,
			&objs,
			repository.Projection(&bson.DocElem{Name: "name", Value: 1}),
		)
		Expect(err).To(BeNil())
		Expect(objs).To(HaveLen(3))
		Expect(objs[0]).To(HaveLen(2))
		Expect(objs[0]).To(HaveKey("_id"))
		Expect(objs[0]).To(HaveKeyWithValue("name", "Snake Eyes"))
		Expect(objs[1]).To(HaveLen(2))
		Expect(objs[1]).To(HaveKey("_id"))
		Expect(objs[1]).To(HaveKeyWithValue("name", "Scarlett"))
		Expect(objs[2]).To(HaveLen(2))
		Expect(objs[2]).To(HaveKey("_id"))
		Expect(objs[2]).To(HaveKeyWithValue("name", "Duke"))
	})
})
