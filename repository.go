package repository

type RepositoryProvider interface {
	GetCollectionName() string
	GetQueryRunner() QueryRunner
	GetDefaultCriteria() interface{}
	GetDefaultSorting() []string
}

type Repository struct {
	collection      string
	runner          QueryRunner
	defaultCriteria interface{}
	defaultSorting  []string
}

type RepositoryConfig struct {
	Collection      string
	QueryRunner     QueryRunner
	DefaultCriteria interface{}
	DefaultSorting  []string
}

func NewRepository(config RepositoryConfig) *Repository {
	return &Repository{
		collection:      config.Collection,
		runner:          config.QueryRunner,
		defaultCriteria: config.DefaultCriteria,
		defaultSorting:  config.DefaultSorting,
	}
}

func (r *Repository) GetCollectionName() string {
	return r.collection
}

func (r *Repository) GetQueryRunner() QueryRunner {
	return r.runner
}

func (r *Repository) GetDefaultCriteria() interface{} {
	return r.defaultCriteria
}

func (r *Repository) GetDefaultSorting() []string {
	return r.defaultSorting
}
