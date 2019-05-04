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

func (r *Repository) CountAndFindAll(dst interface{}, params ...interface{}) (n int, err error) {
	return CountAndFindAll(r, dst, params...)
}

func (r *Repository) Count(params ...interface{}) (n int, err error) {
	return Count(r, params...)
}

func (r *Repository) Create(object interface{}) error {
	return Create(r, object)
}

func (r *Repository) DeleteAll(params ...interface{}) (n int, resultErr error) {
	return DeleteAll(r, params...)
}

func (r *Repository) Delete(params ...interface{}) error {
	return Delete(r, params...)
}

func (r *Repository) FindAll(objects interface{}, params ...interface{}) error {
	return FindAll(r, objects, params...)
}

func (r *Repository) FindByID(id interface{}, dst interface{}, params ...interface{}) error {
	return FindByID(r, id, dst, params...)
}

func (r *Repository) Find(dst interface{}, params ...interface{}) error {
	return Find(r, dst, params...)
}

func (r *Repository) UpdateRawWithCriteria(object interface{}, selector ...interface{}) error {
	return UpdateRawWithCriteria(r, object, selector...)
}

func (r *Repository) UpdateAndFind(id interface{}, dst interface{}, object interface{}, params ...interface{}) error {
	return UpdateAndFind(r, id, dst, object, params...)
}

func (r *Repository) Update(id interface{}, object interface{}) error {
	return Update(r, id, object)
}

func (r *Repository) UpdateRaw(id interface{}, object interface{}) error {
	return UpdateRaw(r, id, object)
}

func (r *Repository) Upsert(object interface{}, selector ...interface{}) error {
	return Upsert(r, object, selector...)
}
