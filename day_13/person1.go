package structure

import "fmt"

type Person struct {
	ID   uint
	name string
	age  int
}

func NewPerson(ID uint, name string, age int) Person {
	return Person{ID, name, age}
}

func PrintPerson() {
	p := NewPerson(1, "James", 48)
	fmt.Println(p)
}
