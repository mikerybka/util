package util

type DB struct {
	FS     FileSystem
	Schema struct {
		Tables map[string]struct {
			Columns map[string]string
		}
	}
}

func (db *DB) CreateTable(name string) error
func (db *DB) AddColumn(table, column, typ string) error
func (db *DB) Insert(name string) error
func (db *DB) Update(name string) error
func (db *DB) Delete(name string) error
func (db *DB) SelectOne(name string) error
func (db *DB) SelectMany(name string) error
