package repository_test

import (
	"github.com/globalsign/mgo/bson"
	"github.com/lab259/repository"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Delete", func() {
	BeforeEach(func() {
		Expect(clearSession()).To(BeNil())
	})

	It("should delete an object (default repository)", func() {
		r := NewRepository()
		tobj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Snake Eyes",
			Age:  33,
		}
		err := r.Create(tobj)
		Expect(err).To(BeNil())
		Expect(r.Delete(repository.WithCriteria(repository.EQ("name", "Snake Eyes")))).To(BeNil())
		objs := make([]testRepObject, 0)
		Expect(r.FindAll(&objs)).To(BeNil())
		Expect(objs).To(BeEmpty())
	})

	It("should delete an object", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		tobj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Snake Eyes",
			Age:  33,
		}
		err := repository.Create(r, tobj)
		Expect(err).To(BeNil())
		Expect(repository.Delete(r, repository.WithCriteria(repository.EQ("name", "Snake Eyes")))).To(BeNil())
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(r, &objs)).To(BeNil())
		Expect(objs).To(BeEmpty())
	})

	It("should delete only the first one of a 2 match", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		insertObjects(r)
		Expect(repository.Delete(r, repository.WithCriteria(
			repository.LT("age", 30),
		))).To(BeNil())
		objs := make([]testRepObject, 0)
		Expect(repository.FindAll(r, &objs)).To(BeNil())
		Expect(objs).To(HaveLen(2))
		Expect(objs[0].Name).To(Equal("Snake Eyes"))
		Expect(objs[1].Name).To(Equal("Duke"))
	})
})
