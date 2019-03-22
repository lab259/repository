package repository_test

import (
	"github.com/lab259/repository"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Operators Evaluation:", func() {
	BeforeEach(func() {
		Expect(clearSession()).To(BeNil())
		connect()
	})

	Context("Operator: $regex", func() {
		It("should find objects with a $regex condition", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.And(
				repository.WithCriteria(
					repository.Regex("name", "Eyes", "i"),
				),
			))).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
		})

		It("should find not exists objects with a $regex", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.Not(
				repository.Regex("name", "Duke", ""),
			))).To(BeNil())
			Expect(objs).To(HaveLen(2))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
			Expect(objs[1].Name).To(Equal("Scarlett"))
		})
	})

	Context("Operator: $text", func() {
		It("should find objects with a $text condition", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			createIndexes(r)
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.Text(
				repository.FindText{
					Search: "eyes",
				},
			))).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
		})

		It("should find objects with a $text condition not matches", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			createIndexes(r)
			insertObjects(r)
			objs := make([]testRepObject, 0)
			Expect(repository.FindAll(r, &objs, repository.Text(
				repository.FindText{
					Search: "Jack",
				},
			))).To(BeNil())
			Expect(objs).To(HaveLen(0))
		})

		It("should fail with a $text condition not created indexes", func() {
			r := &testRepNoDefaultCriteriaNoDefaultSorting{}
			insertObjects(r)
			objs := make([]testRepObject, 0)
			err := repository.FindAll(r, &objs, repository.Text(
				repository.FindText{
					Search: "Duke",
				},
			))
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("text index required for $text query"))
		})

	})

})
