package main

import (
	"fmt"

	"musa.io/lite_database/lib"
)

func main() {
	data := map[string]interface{}{
		"Name": "Musa",
	}

	person := lib.Person{}
	person.Create(data)
	person.Retrive(12)
	person.Update(data)
	person.Delete(12)

	_, err := lib.NewSQLiteDatabase("test.db")

	if err != nil {
		fmt.Println(err)
	}

}
