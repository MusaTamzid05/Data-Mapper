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
	personMapper.Add(lib.Person{Name: "Musa", Age: 27})

	person2 := personMapper.Get(1)
	fmt.Println(person2)
}
