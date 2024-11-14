package loop

import "fmt"

func WhileLoop() {
	var i = 1
	for i <= 25 {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()
}
