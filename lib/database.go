package lib

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// this has bride pattern with the table.go
// this is the implementation hierarchy
type DatabaseInterface interface {
	GetDatabase() interface{}
	GetDatabaseType() string
}

type SQLiteDatabase struct {
	db *gorm.DB
}

func NewSQLiteDatabase(dbPath string) (*SQLiteDatabase, error) {
	var err error
	lite := SQLiteDatabase{}

	lite.db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &lite, nil

}

func (sqlite *SQLiteDatabase) GetDatabase() interface{} {
	return sqlite.db
}

func (sqlite *SQLiteDatabase) GetDatabaseType() string {
	return "sqlite"
}
