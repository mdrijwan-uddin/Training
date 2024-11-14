package main

import (
	loop "Training/day_11"
	set "Training/day_12"
	structure "Training/day_13"
	"fmt"
)

func main() {
	fmt.Println("-----------While loop-------------")
	loop.WhileLoop()
	fmt.Println("------------For loop--------------")
	loop.ForLoop()
	fmt.Println("------------Factorial-------------")
	loop.Factorial(9)
	fmt.Println("--------------Array---------------")
	set.ArraysSum()
	fmt.Println("--------------Slice---------------")
	set.SliceAppend()
	fmt.Println("---------------Map----------------")
	set.MapPrint()
	fmt.Println("-------------Person---------------")
	structure.PrintPerson()
	fmt.Println("----------Person Detail-----------")
	structure.PrintPersonDetail()
	fmt.Println("-------------My Math--------------")
	structure.MyMathMain() 
}
