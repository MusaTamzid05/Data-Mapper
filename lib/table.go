package lib

import (
	"gorm.io/gorm"
)

// data mapper pattern => https://martinfowler.com/eaaCatalog/dataMapper.html

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
	mapper.database.DB.Save(&person)

}

func (mapper *PersonMapper) Delete(person Person) error {
	result := mapper.database.DB.Delete(&person)
	return result.Error

}

func (mapper *PersonMapper) GetAll() []Person {
	persons := []Person{}
	mapper.database.DB.Find(&persons)

	return persons

}
