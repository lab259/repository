package repository_test

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/lab259/repository"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UpdateRawWithCriteria", func() {
	BeforeEach(func() {
		Expect(clearSession()).To(BeNil())
	})

	It("should update an object by one field (default repository)", func() {
		r := NewRepository()
		obj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Snake Eyes",
			Age:  33,
		}
		Expect(r.Create(obj)).To(BeNil())
		Expect(r.UpdateRawWithCriteria(
			bson.M{
				"$set": testRepObject{
					ID:   obj.ID,
					Name: "Scarlett",
					Age:  22,
				},
			}, repository.EQ("name", "Snake Eyes"),
		)).To(BeNil())
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

	It("should update an object by one field", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		obj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Snake Eyes",
			Age:  33,
		}
		Expect(repository.Create(r, obj)).To(BeNil())
		Expect(repository.UpdateRawWithCriteria(r,
			bson.M{
				"$set": testRepObject{
					ID:   obj.ID,
					Name: "Scarlett",
					Age:  22,
				},
			}, repository.EQ("name", "Snake Eyes"),
		)).To(BeNil())
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

	It("should update an object by multiple fields", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		obj := &testRepObject{
			ID:      bson.NewObjectId(),
			Name:    "Snake Eyes",
			Age:     33,
			Status:  true,
			Agility: 10,
		}
		Expect(repository.Create(r, obj)).To(BeNil())
		Expect(repository.UpdateRawWithCriteria(r,
			bson.M{
				"$set": &testRepObject{
					ID:       obj.ID,
					Name:     "Scarlett",
					Age:      22,
					Agility:  5,
					Strength: 50,
				},
			},
			repository.EQ("age", 33),
			repository.EQ("status", true),
			repository.EQ("agility", 10),
		)).To(BeNil())
		Expect(defaultQueryRunner.RunWithDB(func(db *mgo.Database) error {
			c := db.C(r.GetCollectionName())
			objs := make([]testRepObject, 0)
			Expect(c.Find(nil).All(&objs)).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Scarlett"))
			Expect(objs[0].Age).To(Equal(22))
			Expect(objs[0].Agility).To(Equal(5))
			Expect(objs[0].Strength).To(Equal(50))
			return nil
		})).To(BeNil())
	})

	It("should update multiple objects by one field", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		obj := &testRepObject{
			ID:      bson.NewObjectId(),
			Name:    "Snake Eyes",
			Age:     33,
			Agility: 10,
		}
		Expect(repository.Create(r, obj)).To(BeNil())
		obj1 := &testRepObject{
			ID:      bson.NewObjectId(),
			Name:    "Deadpool",
			Age:     25,
			Agility: 10,
		}
		Expect(repository.Create(r, obj1)).To(BeNil())
		Expect(repository.UpdateRawWithCriteria(r,
			bson.M{
				"$set": &testRepObject{
					Agility: 99,
				},
			}, repository.EQ("agility", 10),
		)).To(BeNil())
		Expect(defaultQueryRunner.RunWithDB(func(db *mgo.Database) error {
			c := db.C(r.GetCollectionName())
			objs := make([]testRepObject, 0)
			Expect(c.Find(nil).All(&objs)).To(BeNil())
			Expect(objs).To(HaveLen(2))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
			Expect(objs[0].Age).To(Equal(33))
			Expect(objs[0].Agility).To(Equal(99))

			Expect(objs[1].Name).To(Equal("Deadpool"))
			Expect(objs[1].Age).To(Equal(25))
			Expect(objs[1].Agility).To(Equal(99))
			return nil
		})).To(BeNil())
	})

	It("should update multiple objects by multiple fields", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		obj := &testRepObject{
			ID:       bson.NewObjectId(),
			Name:     "Snake Eyes",
			Age:      33,
			Agility:  10,
			Strength: 40,
		}
		Expect(repository.Create(r, obj)).To(BeNil())
		obj1 := &testRepObject{
			ID:       bson.NewObjectId(),
			Name:     "Deadpool",
			Age:      25,
			Agility:  10,
			Strength: 40,
		}
		Expect(repository.Create(r, obj1)).To(BeNil())
		Expect(repository.UpdateRawWithCriteria(r,
			bson.M{
				"$set": &testRepObject{
					Agility:  99,
					Strength: 50,
				},
			},
			repository.EQ("agility", 10),
			repository.EQ("strength", 40),
		)).To(BeNil())
		Expect(defaultQueryRunner.RunWithDB(func(db *mgo.Database) error {
			c := db.C(r.GetCollectionName())
			objs := make([]testRepObject, 0)
			Expect(c.Find(nil).All(&objs)).To(BeNil())
			Expect(objs).To(HaveLen(2))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
			Expect(objs[0].Age).To(Equal(33))
			Expect(objs[0].Agility).To(Equal(99))
			Expect(objs[0].Strength).To(Equal(50))

			Expect(objs[1].Name).To(Equal("Deadpool"))
			Expect(objs[1].Age).To(Equal(25))
			Expect(objs[1].Agility).To(Equal(99))
			Expect(objs[1].Strength).To(Equal(50))
			return nil
		})).To(BeNil())
	})

	It("should fail updating an non existent object", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		Expect(repository.UpdateRawWithCriteria(r,
			bson.M{
				"$set": testRepObject{
					Name: "Scarlett",
					Age:  22,
				},
			}, repository.EQ("age", 22))).To(Equal(mgo.ErrNotFound))
	})

	It("should update an object using no helper", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		obj := &testRepObject{
			ID:     bson.NewObjectId(),
			Name:   "Snake Eyes",
			Age:    33,
			Status: true,
		}
		Expect(repository.Create(r, obj)).To(BeNil())
		Expect(repository.UpdateRawWithCriteria(r,
			map[string]interface{}{
				"$inc": map[string]interface{}{
					"age": 1,
				},
				"$set": map[string]interface{}{
					"status": false,
				},
			},
			map[string]interface{}{
				"_id": obj.ID,
			})).To(BeNil())
		Expect(defaultQueryRunner.RunWithDB(func(db *mgo.Database) error {
			c := db.C(r.GetCollectionName())
			objs := make([]testRepObject, 0)
			Expect(c.Find(nil).All(&objs)).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
			Expect(objs[0].Age).To(Equal(34))
			Expect(objs[0].Status).To(BeFalse())
			return nil
		})).To(BeNil())
	})
})
