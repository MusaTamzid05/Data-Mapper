package lib

import "fmt"

// this has bride pattern with database.go
// this is the implementation part of the pattern
// this is the anbstruction hierarchy

type Table interface {
	Init(database DatabaseInterface)
	Create(row map[string]interface{})
	Retrive(id int)
	Update(row map[string]interface{})
	Delete(id int)
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) Init(database DatabaseInterface) {

}

func (p *Person) Create(row map[string]interface{}) {
	fmt.Println("Creating a person")
}

func (p *Person) Retrive(id int) {
	fmt.Println("Retriving a person")

}

func (p *Person) Update(row map[string]interface{}) {
	fmt.Println("Update a person")

}

func (p *Person) Delete(id int) {
	fmt.Println("Delete a person")

}
