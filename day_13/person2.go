package structure

import (
	"fmt"
	"strconv"
)

type Person2 struct {
	ID   uint
	name string
	age  int
}

func NewPerson2(ID uint, name string, age int) Person2 {
	return Person2{ID, name, age}
}

func (p Person2) String() string {
	return "PersonID: " + strconv.Itoa(int(p.ID)) + "\tName: " + p.name + "\tAge: " + strconv.Itoa(p.age)
}

func PrintPersonDetail() {
	p := NewPerson2(2, "Moses", 33)
	fmt.Println(p)
}
