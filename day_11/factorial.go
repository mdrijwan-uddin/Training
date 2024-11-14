package loop

import "fmt"

func Factorial(n int) {
	var i = 1
	var sum int
	for i <= n {
		sum += i
		i++
	}
	fmt.Printf("Factorial 9: ")
	fmt.Println(sum)
}
