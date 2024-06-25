package util

type Index map[string]Set

func (i Index) Get(key string) Set {
	return i[key]
}

func (i Index) Add(k, v string) {
	i[k].Add(v)
}

func (i Index) Remove(k, v string) {
	i[k].Remove(v)
}
