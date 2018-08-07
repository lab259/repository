package repository_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/mgo.v2/bson"

	"."
	"./queries"
	"gopkg.in/mgo.v2"
)

var _ = Describe("FindByID", func() {
	insertObjects := func(r repository.Repository) (bson.ObjectId, bson.ObjectId, bson.ObjectId) {
		objId1 := bson.NewObjectId()
		objId2 := bson.NewObjectId()
		objId3 := bson.NewObjectId()
		Expect(repository.Create(r, &testRepObject{
			ID:   objId1,
			Name: "Snake Eyes",
			Age:  33,
		})).To(BeNil())
		Expect(repository.Create(r, &testRepObject{
			ID:   objId2,
			Name: "Scarlett",
			Age:  22,
		})).To(BeNil())
		Expect(repository.Create(r, &testRepObject{
			ID:   objId3,
			Name: "Duke",
			Age:  22,
		})).To(BeNil())

		return objId1, objId2, objId3
	}

	BeforeEach(func() {
		Expect(clearSession()).To(BeNil())
	})

	It("should find an object by id", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		objid1, _, _ := insertObjects(r)
		var obj testRepObject
		Expect(repository.FindByID(r, objid1, &obj)).To(BeNil())
		Expect(obj.Name).To(Equal("Snake Eyes"))
		Expect(obj.Age).To(Equal(33))
	})

	It("should not find an object by id", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		objid1 := bson.NewObjectId()
		var obj testRepObject
		Expect(repository.FindByID(r, objid1, &obj)).To(Equal(mgo.ErrNotFound))
	})

	It("should find an object by id and another criteria", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		objid1, _, _ := insertObjects(r)
		var obj testRepObject
		Expect(repository.FindByID(r, objid1, &obj, queries.WithCriteria(
			queries.GT("age", 30),
		))).To(BeNil())
		Expect(obj.Name).To(Equal("Snake Eyes"))
		Expect(obj.Age).To(Equal(33))
	})

	It("should find an object by id and another criteria", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		objid1, _, _ := insertObjects(r)
		var obj testRepObject
		Expect(repository.FindByID(r, objid1, &obj, queries.EQ("age", 33))).To(BeNil())
		Expect(obj.Name).To(Equal("Snake Eyes"))
		Expect(obj.Age).To(Equal(33))
	})

	It("should find an object by id and another criteria", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		objid1, _, _ := insertObjects(r)
		var obj testRepObject
		Expect(repository.FindByID(r, objid1, &obj, queries.WithCriteria(
			queries.LT("age", 30),
		))).To(Equal(mgo.ErrNotFound))
	})

	It("should fail finding an object with a criteria with error", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		objid1, _, _ := insertObjects(r)
		var obj testRepObject
		repository.FindByID(r, objid1, &obj, queries.WithCriteria(
			&erroneousCondition{},
		))
	})

	It("should fail finding an object with a erroneous criteria", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		objid1, _, _ := insertObjects(r)
		var obj testRepObject
		repository.FindByID(r, objid1, &obj, &erroneousCondition{})
	})

	It("should fail finding an object with a not supported filter type", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		objid1, _, _ := insertObjects(r)
		var obj testRepObject
		err := repository.FindByID(r, objid1, &obj, bson.M{})
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(ContainSubstring("data type not supported"))
	})

	It("should fail finding an object with a erroneous QueryModifier", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		objid1, _, _ := insertObjects(r)
		var obj testRepObject
		err := repository.FindByID(r, objid1, &obj, &erroneousQueryModifier{})
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(ContainSubstring("forced error"))
	})

	It("should fail finding an object with a erroneous QueryModifier", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		objid1, _, _ := insertObjects(r)
		var obj testRepObject
		err := repository.FindByID(r, objid1, &obj, queries.WithCriteria(
			&erroneousCondition{},
		))
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(ContainSubstring("forced error"))
	})
})
