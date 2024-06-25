package util

import (
	"fmt"
	"reflect"
)

func NewTable[T any]() *Table[T] {
	t := &Table[T]{
		Rows:    make(map[string]T),
		Indexes: make(map[string]Index),
	}

	var row T
	typ := reflect.TypeOf(row)
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		panic("type not struct")
	}
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		t.AddIndex(f.Name)
	}

	return t
}

type Table[T any] struct {
	Rows    map[string]T
	Indexes map[string]Index
}

func (t *Table[T]) AddIndex(fieldID string) error {
	// Make sure the field exists and is a string type
	var row T
	typ := reflect.TypeOf(row)
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		panic("type not struct")
	}
	f, ok := typ.FieldByName(fieldID)
	if !ok {
		return fmt.Errorf("no field %s", fieldID)
	}
	if f.Type.Kind() != reflect.String {
		return fmt.Errorf("field %s not string", fieldID)
	}

	// Add the index to the map
	t.Indexes[fieldID] = make(Index)

	return nil
}

func (t *Table[T]) Find(id string) (T, bool) {
	v, ok := t.Rows[id]
	return v, ok
}

func (t *Table[T]) IDs() Set {
	s := Set{}
	for k := range t.Rows {
		s.Add(k)
	}
	return s
}

func (t *Table[T]) Select(match ...Pair[string, string]) map[string]T {
	ids := t.IDs()
	for _, cond := range match {
		index := t.Indexes[cond.K]
		set := index.Get(cond.V)
		ids = ids.Intersection(set)
	}

	results := map[string]T{}

	for _, id := range ids.List() {
		results[id] = t.Rows[id]
	}

	return results
}

func (t *Table[T]) Insert(v T) {
	// Generate a new ID.
	id := RandomID()

	// Update the data.
	t.Rows[id] = v

	// Update the indexes.
	for fieldID, index := range t.Indexes {
		value := reflect.ValueOf(v).FieldByName(fieldID).Interface().(string)
		index.Add(value, id)
	}
}

func (t *Table[T]) Update(id string, v T) {
	old := t.Rows[id]

	// Update the indexes.
	for fieldID, index := range t.Indexes {
		oldValue := reflect.ValueOf(old).FieldByName(fieldID).Interface().(string)
		newValue := reflect.ValueOf(v).FieldByName(fieldID).Interface().(string)
		if oldValue != newValue {
			index.Remove(oldValue, id)
			index.Add(newValue, id)
		}
	}

	// Update the data.
	t.Rows[id] = v
}

func (t *Table[T]) Delete(id string) {
	old := t.Rows[id]

	// Update the indexes.
	for fieldID, index := range t.Indexes {
		oldValue := reflect.ValueOf(old).FieldByName(fieldID).Interface().(string)
		index.Remove(oldValue, id)
	}

	// Update the data.
	delete(t.Rows, id)
}
