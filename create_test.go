package repository_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"

	"."
)

var _ = Describe("Create", func() {
	BeforeEach(func() {
		Expect(clearSession()).To(BeNil())
	})

	It("should create an object", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		err := repository.Create(r, testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Snake Eyes",
			Age:  33,
		})
		Expect(err).To(BeNil())
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
})
