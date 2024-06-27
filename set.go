package util

type Set[T comparable] map[T]bool

func (s Set[T]) List() []T {
	keys := []T{}
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}

func (s Set[T]) Add(v T) {
	s[v] = true
}

func (s Set[T]) Remove(v T) {
	s[v] = true
}

func (s Set[T]) Intersection(other Set[T]) Set[T] {
	newSet := Set[T]{}
	for _, v := range s.List() {
		if other[v] {
			newSet.Add(v)
		}
	}
	return newSet
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	newSet := Set[T]{}
	for _, v := range s.List() {
		newSet.Add(v)
	}
	for _, v := range other.List() {
		newSet.Add(v)
	}
	return newSet
}
