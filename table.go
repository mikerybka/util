package util

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"sync"
)

type Table[T any] struct {
	DataDir    string
	IndexDir   string
	Indexes    Set[string]
	postLock   *sync.Mutex
	indexLocks map[string]*sync.Mutex
}

func (t *Table[T]) ListAll() []string {
	entries, _ := os.ReadDir(t.DataDir)
	ids := []string{}
	for _, e := range entries {
		ids = append(ids, e.Name())
	}
	return ids
}

func (t *Table[T]) Get(id string) (*T, error) {
	path := filepath.Join(t.DataDir, id)
	v := new(T)
	err := ReadJSONFile(path, v)
	return v, err
}

func (t *Table[T]) FindIDsBy(k, v string) (Set[string], error) {
	if !t.Indexes.Has(k) {
		return nil, fmt.Errorf("field %s not indexed", k)
	}
	path := filepath.Join(t.IndexDir, k, v)
	var index Set[string]
	ReadJSONFile(path, index)
	return index, nil
}

func (t *Table[T]) Delete(id string) error {
	// Remove indexes
	oldValue, _ := t.Get(id)
	for index := range t.Indexes {
		val := reflect.ValueOf(oldValue)
		fieldValue := val.Elem().FieldByName(index)
		path := filepath.Join(t.IndexDir, index, fmt.Sprintf("%s", fieldValue.Interface()))
		t.indexLocks[index].Lock()
		var i Set[string]
		err := ReadJSONFile(path, i)
		if err != nil {
			t.indexLocks[index].Unlock()
			panic(err)
		}
		i.Remove(id)
		err = WriteJSONFile(path, i)
		t.indexLocks[index].Unlock()
		if err != nil {
			return fmt.Errorf("writing index %s: %s", index, err)
		}
	}

	// Remove data
	path := filepath.Join(t.DataDir, id)
	return os.Remove(path)
}

func (t *Table[T]) Set(id string, v *T) error {
	// Remove old indexes
	oldValue, _ := t.Get(id)
	for index := range t.Indexes {
		val := reflect.ValueOf(oldValue)
		fieldValue := val.Elem().FieldByName(index)
		path := filepath.Join(t.IndexDir, index, fmt.Sprintf("%s", fieldValue.Interface()))
		t.indexLocks[index].Lock()
		var i Set[string]
		err := ReadJSONFile(path, i)
		if err != nil {
			t.indexLocks[index].Unlock()
			panic(err)
		}
		i.Remove(id)
		err = WriteJSONFile(path, i)
		t.indexLocks[index].Unlock()
		if err != nil {
			return fmt.Errorf("writing index %s: %s", index, err)
		}
	}

	// Write new indexes
	for index := range t.Indexes {
		val := reflect.ValueOf(v)
		fieldValue := val.Elem().FieldByName(index)
		if fieldValue.Type().String() != "string" {
			panic("only string fields can be indexed")
		}
		path := filepath.Join(t.IndexDir, index, fmt.Sprintf("%s", fieldValue.Interface()))
		t.indexLocks[index].Lock()
		var i Set[string]
		err := ReadJSONFile(path, i)
		if err != nil {
			t.indexLocks[index].Unlock()
			panic(err)
		}
		i.Add(id)
		err = WriteJSONFile(path, i)
		t.indexLocks[index].Unlock()
		if err != nil {
			return fmt.Errorf("writing index %s: %s", index, err)
		}
	}

	// Write data
	path := filepath.Join(t.DataDir, id)
	err := WriteJSONFile(path, v)
	if err != nil {
		return fmt.Errorf("writing data: %s", err)
	}

	return nil
}

func (t *Table[T]) Len() int {
	entries, _ := os.ReadDir(t.DataDir)
	return len(entries)
}

func (t *Table[T]) Post(v *T) error {
	t.postLock.Lock()
	defer t.postLock.Unlock()
	id := strconv.Itoa(t.Len())
	return t.Set(id, v)
}
