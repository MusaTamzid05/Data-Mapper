package main

import (
	"fmt"

	"musa.io/lite_database/lib"
)

func main() {

	db, err := lib.NewSQLiteDatabase("test.db")

	if err != nil {
		fmt.Println(err)
	}

	db.Migrate()
	personMapper := lib.NewPersonMapper(db)

	/*
		personMapper.Add(lib.Person{Name: "Musa", Age: 27})

		person2 := personMapper.Get(1)
		fmt.Println(person2)

		person2.Name = "Tamzid"
		personMapper.Update(person2)

		delPerson := personMapper.Get(7)
		personMapper.Delete(delPerson)
	*/

	persons := personMapper.GetAll()

	for _, person := range persons {
		fmt.Println(person)
	}

	fmt.Println("===========================")

	delPerson := personMapper.Get(9)
	err = personMapper.Delete(delPerson)

	if err != nil {
		fmt.Println(err)
	}

	persons = personMapper.GetAll()

	for _, person := range persons {
		fmt.Println(person)
	}

}
