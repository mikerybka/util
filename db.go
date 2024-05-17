package util

type DB struct {
	FS     FileSystem
	Schema struct {
		Tables map[string]struct {
			Columns map[string]string
		}
	}
}

func (db *DB) CreateTable(name string) error {
	panic("not implemented")
}
func (db *DB) AddColumn(table, column, typ string) error {
	panic("not implemented")
}
func (db *DB) Insert(name string) error {
	panic("not implemented")
}
func (db *DB) Update(name string) error {
	panic("not implemented")
}
func (db *DB) Delete(name string) error {
	panic("not implemented")
}
func (db *DB) SelectOne(name string) error {
	panic("not implemented")
}
func (db *DB) SelectMany(name string) error {
	panic("not implemented")
}
