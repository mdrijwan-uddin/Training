package structure

import (
	"fmt"
)

type MyMath interface {
	Add() float64
	Subtract() float64
	Multiply() float64
	Divide() float64
}

type MyNum struct {
	num1, num2 float64
}

func (m MyNum) Add() float64 {
	return m.num1 + m.num2
}

func (m MyNum) Subtract() float64 {
	return m.num1 - m.num2
}

func (m MyNum) Multiply() float64 {
	return m.num1 * m.num2
}

func (m MyNum) Divide() float64 {
	return m.num1 / m.num2
}

func PrintMyMath(m MyMath) {
	fmt.Println(m)
	fmt.Println(m.Add())
	fmt.Println(m.Subtract())
	fmt.Println(m.Multiply())
	fmt.Println(m.Divide())
}

func MyMathMain() {
	m := MyNum{10, 5}

	PrintMyMath(m)
}
