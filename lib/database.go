package lib

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteDatabase struct {
	DB *gorm.DB
}

func NewSQLiteDatabase(dbPath string) (*SQLiteDatabase, error) {
	var err error
	lite := SQLiteDatabase{}

	lite.DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &lite, nil

}

func (lite *SQLiteDatabase) Migrate() {
	lite.DB.AutoMigrate(&Person{})

}
