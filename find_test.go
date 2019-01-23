package repository_test

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/lab259/repository"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Find", func() {
	BeforeEach(func() {
		Expect(clearSession()).To(BeNil())
	})

	It("should find an object by name", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		tobj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Snake Eyes",
			Age:  33,
		}
		err := repository.Create(r, tobj)
		Expect(err).To(BeNil())
		var obj testRepObject
		Expect(repository.Find(r, &obj, repository.WithCriteria(repository.EQ("name", "Snake Eyes")))).To(BeNil())
	})

	It("should not find an object by name", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		tobj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Snake Eyes",
			Age:  33,
		}
		err := repository.Create(r, tobj)
		Expect(err).To(BeNil())
		var obj testRepObject
		Expect(repository.Find(r, &obj, repository.WithCriteria(repository.EQ("name", "Snake Eyes 1")))).To(Equal(mgo.ErrNotFound))
	})

	It("should combine criteria to find an object", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		tobj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Snake Eyes",
			Age:  33,
		}
		err := repository.Create(r, tobj)
		Expect(err).To(BeNil())
		var obj testRepObject
		Expect(repository.Find(r, &obj, repository.WithCriteria(
			repository.And(
				repository.EQ("name", "Snake Eyes"),
				repository.GT("age", 30),
			),
		))).To(BeNil())
		Expect(obj.Name).To(Equal("Snake Eyes"))
		Expect(obj.Age).To(Equal(33))
	})

	It("should not find an object with greater than operator", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		tobj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Snake Eyes",
			Age:  33,
		}
		err := repository.Create(r, tobj)
		Expect(err).To(BeNil())
		var obj testRepObject
		Expect(repository.Find(r, &obj, repository.WithCriteria(
			repository.GT("age", 33),
		))).To(Equal(mgo.ErrNotFound))
	})

	It("should find an object with greater than or equal operator", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		tobj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Snake Eyes",
			Age:  33,
		}
		err := repository.Create(r, tobj)
		Expect(err).To(BeNil())
		var obj testRepObject
		Expect(repository.Find(r, &obj, repository.WithCriteria(
			repository.And(
				repository.EQ("name", "Snake Eyes"),
				repository.GTE("age", 33),
			),
		))).To(BeNil())
		Expect(obj.Name).To(Equal("Snake Eyes"))
		Expect(obj.Age).To(Equal(33))
	})

	It("should not find an object with greater than or equal operator", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		tobj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Snake Eyes",
			Age:  33,
		}
		err := repository.Create(r, tobj)
		Expect(err).To(BeNil())
		var obj testRepObject
		Expect(repository.Find(r, &obj, repository.WithCriteria(
			repository.GTE("age", 34),
		))).To(Equal(mgo.ErrNotFound))
	})

	It("should find an object with less than or equal operator", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		tobj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Snake Eyes",
			Age:  33,
		}
		err := repository.Create(r, tobj)
		Expect(err).To(BeNil())
		var obj testRepObject
		Expect(repository.Find(r, &obj, repository.WithCriteria(
			repository.LTE("age", 33),
		))).To(BeNil())
		Expect(obj.Name).To(Equal("Snake Eyes"))
		Expect(obj.Age).To(Equal(33))
	})

	It("should not find an object with less than or equal operator", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		tobj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Snake Eyes",
			Age:  33,
		}
		err := repository.Create(r, tobj)
		Expect(err).To(BeNil())
		var obj testRepObject
		Expect(repository.Find(r, &obj, repository.WithCriteria(
			repository.LTE("age", 32),
		))).To(Equal(mgo.ErrNotFound))
	})

	It("should find an object with exists operator", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		tobj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Snake Eyes",
			Age:  33,
		}
		err := repository.Create(r, tobj)
		Expect(err).To(BeNil())
		var obj testRepObject
		Expect(repository.Find(r, &obj, repository.WithCriteria(
			repository.Exists("age", true),
		))).To(BeNil())
		Expect(obj.Name).To(Equal("Snake Eyes"))
		Expect(obj.Age).To(Equal(33))
	})

	It("should not find an object with the exists operator", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		tobj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Snake Eyes",
			Age:  33,
		}
		err := repository.Create(r, tobj)
		Expect(err).To(BeNil())
		var obj testRepObject
		Expect(repository.Find(r, &obj, repository.WithCriteria(
			repository.Exists("age", false),
		))).To(Equal(mgo.ErrNotFound))
	})
})
