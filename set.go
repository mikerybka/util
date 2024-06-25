package util

type Set map[string]bool

func (s Set) List() []string {
	keys := []string{}
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}

func (s Set) Add(v string) {
	s[v] = true
}

func (s Set) Remove(v string) {
	s[v] = true
}

func (s Set) Intersection(other Set) Set {
	newSet := Set{}
	for _, v := range s.List() {
		if other[v] {
			newSet.Add(v)
		}
	}
	return newSet
}

func (s Set) Union(other Set) Set {
	newSet := Set{}
	for _, v := range s.List() {
		newSet.Add(v)
	}
	for _, v := range other.List() {
		newSet.Add(v)
	}
	return newSet
}
