package lib

import (
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name string
	Age  int
}

type PersonMapper struct {
	database *SQLiteDatabase
}

func NewPersonMapper(database *SQLiteDatabase) *PersonMapper {
	mapper := PersonMapper{database: database}
	return &mapper

}

func (mapper *PersonMapper) Get(id uint) Person {
	person := Person{}
	mapper.database.DB.First(&person)
	return person

}

func (mapper *PersonMapper) Add(person Person) error {
	result := mapper.database.DB.Create(&person)
	return result.Error

}

func (mapper *PersonMapper) Update(person Person) {

}

func (mapper *PersonMapper) Delete(id uint) {

}

func (mapper *PersonMapper) GetAll(id uint) []Person {
	persons := []Person{}
	return persons

}
