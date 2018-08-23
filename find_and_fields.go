package repository

func FindAndFilters(r Repository, dst interface{}, filters []interface{}, params ...interface{}) (n int, err error) {
	return CountAndFindAll(r, dst,
		append([]interface{}{
			&Criteria{
				Conditions: filters,
			},
		}, params...)...)
}
