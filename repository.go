package repository

type Repository interface {
	GetCollectionName() string
	GetQueryRunner() QueryRunner
	GetDefaultCriteria() interface{}
	GetDefaultSorting() []string
}
