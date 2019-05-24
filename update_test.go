package repository_test

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/lab259/repository"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Update", func() {
	BeforeEach(func() {
		Expect(clearSession()).To(BeNil())
	})

	It("should update an object (default repository)", func() {
		r := NewRepository()
		obj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Snake Eyes",
			Age:  33,
		}
		Expect(r.Create(obj)).To(BeNil())
		Expect(r.Update(obj.ID, testRepObject{
			ID:   obj.ID,
			Name: "Scarlett",
			Age:  22,
		})).To(BeNil())
		Expect(defaultQueryRunner.RunWithDB(func(db *mgo.Database) error {
			c := db.C(r.GetCollectionName())
			objs := make([]testRepObject, 0)
			Expect(c.Find(nil).All(&objs)).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Scarlett"))
			Expect(objs[0].Age).To(Equal(22))
			return nil
		})).To(BeNil())
	})
	It("should update an object", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		obj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Snake Eyes",
			Age:  33,
		}
		Expect(repository.Create(r, obj)).To(BeNil())
		Expect(repository.Update(r, obj.ID, testRepObject{
			ID:   obj.ID,
			Name: "Scarlett",
			Age:  22,
		})).To(BeNil())
		Expect(defaultQueryRunner.RunWithDB(func(db *mgo.Database) error {
			c := db.C(r.GetCollectionName())
			objs := make([]testRepObject, 0)
			Expect(c.Find(nil).All(&objs)).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Scarlett"))
			Expect(objs[0].Age).To(Equal(22))
			return nil
		})).To(BeNil())
	})

	It("should update partially an object", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		obj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Snake Eyes",
			Age:  33,
		}
		Expect(repository.Create(r, obj)).To(BeNil())
		Expect(repository.Update(r, obj.ID, testRepObject{
			Age: 22,
		})).To(BeNil())
		Expect(defaultQueryRunner.RunWithDB(func(db *mgo.Database) error {
			c := db.C(r.GetCollectionName())
			objs := make([]testRepObject, 0)
			Expect(c.Find(nil).All(&objs)).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
			Expect(objs[0].Age).To(Equal(22))
			return nil
		})).To(BeNil())
	})

	It("should update and find an object (default repository)", func() {
		r := NewRepository()
		obj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Snake Eyes",
			Age:  33,
		}

		var updatedObj testRepObject
		Expect(r.Create(obj)).To(BeNil())
		Expect(r.UpdateAndFind(obj.ID, &updatedObj, testRepObject{
			Age: 22,
		})).To(BeNil())
		Expect(defaultQueryRunner.RunWithDB(func(db *mgo.Database) error {
			c := db.C(r.GetCollectionName())
			objs := make([]testRepObject, 0)
			Expect(c.Find(nil).All(&objs)).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
			Expect(objs[0].Age).To(Equal(22))
			return nil
		})).To(BeNil())

		Expect(updatedObj.ID).To(Equal(obj.ID))
		Expect(updatedObj.Age).To(Equal(22))
		Expect(updatedObj.Name).To(Equal("Snake Eyes"))
	})

	It("should update and find an object", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		obj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Snake Eyes",
			Age:  33,
		}

		var updatedObj testRepObject
		Expect(repository.Create(r, obj)).To(BeNil())
		Expect(repository.UpdateAndFind(r, obj.ID, &updatedObj, testRepObject{
			Age: 22,
		})).To(BeNil())
		Expect(defaultQueryRunner.RunWithDB(func(db *mgo.Database) error {
			c := db.C(r.GetCollectionName())
			objs := make([]testRepObject, 0)
			Expect(c.Find(nil).All(&objs)).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
			Expect(objs[0].Age).To(Equal(22))
			return nil
		})).To(BeNil())

		Expect(updatedObj.ID).To(Equal(obj.ID))
		Expect(updatedObj.Age).To(Equal(22))
		Expect(updatedObj.Name).To(Equal("Snake Eyes"))
	})

	It("should fail to update and find an object with params", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		obj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Chico Bento",
			Age:  21,
		}

		var updatedObj testRepObject
		Expect(repository.Create(r, obj)).To(BeNil())

		err := repository.UpdateAndFind(r, obj.ID, &updatedObj, testRepObject{
			Name: "Betinho Jr",
			Age:  22,
		}, repository.LT("age", 20))

		Expect(err).To(HaveOccurred())
		Expect(err).To(Equal(mgo.ErrNotFound))

		Expect(defaultQueryRunner.RunWithDB(func(db *mgo.Database) error {
			c := db.C(r.GetCollectionName())
			objs := make([]testRepObject, 0)
			Expect(c.Find(nil).All(&objs)).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Chico Bento"))
			Expect(objs[0].Age).To(Equal(21))
			return nil
		})).To(BeNil())
	})

	It("should update and find an object with params", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		obj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Chico Bento",
			Age:  21,
		}

		var updatedObj testRepObject
		Expect(repository.Create(r, obj)).To(BeNil())

		err := repository.UpdateAndFind(r, obj.ID, &updatedObj, testRepObject{
			Name: "Betinho Jr",
			Age:  22,
		}, repository.GT("age", 20))

		Expect(err).To(BeNil())

		Expect(defaultQueryRunner.RunWithDB(func(db *mgo.Database) error {
			c := db.C(r.GetCollectionName())
			objs := make([]testRepObject, 0)
			Expect(c.Find(nil).All(&objs)).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Betinho Jr"))
			Expect(objs[0].Age).To(Equal(22))
			return nil
		})).To(BeNil())
	})

	It("should fail to update and find an object with params invalid", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		obj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Betinho Jr",
			Age:  21,
		}

		var updatedObj testRepObject
		Expect(repository.Create(r, obj)).To(BeNil())

		err := repository.UpdateAndFind(r, obj.ID, &updatedObj, testRepObject{
			Age: 22,
		}, bson.M{"field": "value"})

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("data type not supported:"))
	})

	It("should fail updating an non existent object", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		Expect(repository.Update(r, bson.NewObjectId(), testRepObject{
			Name: "Scarlett",
			Age:  22,
		})).To(Equal(mgo.ErrNotFound))
	})

	It("should update an object using no helper (default repository)", func() {
		r := NewRepository()
		obj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Snake Eyes",
			Age:  33,
		}
		Expect(r.Create(obj)).To(BeNil())
		Expect(r.UpdateRaw(obj.ID, map[string]interface{}{
			"$inc": map[string]interface{}{
				"age": 1,
			},
		})).To(BeNil())
		Expect(defaultQueryRunner.RunWithDB(func(db *mgo.Database) error {
			c := db.C(r.GetCollectionName())
			objs := make([]testRepObject, 0)
			Expect(c.Find(nil).All(&objs)).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
			Expect(objs[0].Age).To(Equal(34))
			return nil
		})).To(BeNil())
	})

	It("should update an object using no helper", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		obj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Snake Eyes",
			Age:  33,
		}
		Expect(repository.Create(r, obj)).To(BeNil())
		Expect(repository.UpdateRaw(r, obj.ID, map[string]interface{}{
			"$inc": map[string]interface{}{
				"age": 1,
			},
		})).To(BeNil())
		Expect(defaultQueryRunner.RunWithDB(func(db *mgo.Database) error {
			c := db.C(r.GetCollectionName())
			objs := make([]testRepObject, 0)
			Expect(c.Find(nil).All(&objs)).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
			Expect(objs[0].Age).To(Equal(34))
			return nil
		})).To(BeNil())
	})
})
