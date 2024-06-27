package util

type Index map[string]Set[string]

func (i Index) Get(key string) Set[string] {
	return i[key]
}

func (i Index) Add(k, v string) {
	i[k].Add(v)
}

func (i Index) Remove(k, v string) {
	i[k].Remove(v)
}
