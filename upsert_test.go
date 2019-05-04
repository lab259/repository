package repository_test

import (
	"github.com/globalsign/mgo"
	"github.com/lab259/repository"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Upsert", func() {
	BeforeEach(func() {
		Expect(clearSession()).To(BeNil())
	})

	It("should create an object by one field (default repository)", func() {
		r := NewRepository()
		obj := &testRepObject{
			Name: "Snake Eyes",
			Age:  33,
		}

		Expect(r.Upsert(obj)).To(BeNil())
		Expect(defaultQueryRunner.RunWithDB(func(db *mgo.Database) error {
			c := db.C(r.GetCollectionName())
			objs := make([]testRepObject, 0)
			Expect(c.Find(nil).All(&objs)).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
			Expect(objs[0].Age).To(Equal(33))
			return nil
		})).To(BeNil())
	})

	It("should create an object by one field", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		obj := &testRepObject{
			Name: "Snake Eyes",
			Age:  33,
		}

		Expect(repository.Upsert(r, obj)).To(BeNil())
		Expect(defaultQueryRunner.RunWithDB(func(db *mgo.Database) error {
			c := db.C(r.GetCollectionName())
			objs := make([]testRepObject, 0)
			Expect(c.Find(nil).All(&objs)).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
			Expect(objs[0].Age).To(Equal(33))
			return nil
		})).To(BeNil())
	})

	It("should update an object by one field", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		obj := &testRepObject{
			Name: "Duke",
			Age:  33,
		}
		Expect(repository.Create(r, obj)).To(BeNil())

		obj = &testRepObject{
			Name: "Chico Bento",
			Age:  22,
		}
		Expect(repository.Upsert(r, obj)).To(BeNil())
		Expect(defaultQueryRunner.RunWithDB(func(db *mgo.Database) error {
			c := db.C(r.GetCollectionName())
			objs := make([]testRepObject, 0)
			Expect(c.Find(nil).All(&objs)).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Chico Bento"))
			Expect(objs[0].Age).To(Equal(22))
			return nil
		})).To(BeNil())
	})
})
