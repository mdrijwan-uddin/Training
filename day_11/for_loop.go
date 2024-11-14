package loop

import "fmt"

func ForLoop() {
	for i := 0; i < 25; i++ {
		fmt.Printf("%d ", i + 1)
	}
	fmt.Println()
}
