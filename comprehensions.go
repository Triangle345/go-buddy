package comprehensions

func Map[T, U any](collection []T, transform func(T) U) []U {
	result := make([]U, len(collection))
	for i, item := range collection {
		result[i] = transform(item)
	}
	return result
}

func Reduce[T, U any](collection []T, initial U, combine func(U, T) U) U {
	accumulator := initial
	for _, item := range collection {
		accumulator = combine(accumulator, item)
	}
	return accumulator
}

func Filter[T any](collection []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, item := range collection {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func Find[T any](collection []T, predicate func(T) bool) (T, bool) {
	filtered := Filter(collection, predicate)
	if len(filtered) > 0 {
		return filtered[0], true
	}
	var empty T
	return empty, false
}
