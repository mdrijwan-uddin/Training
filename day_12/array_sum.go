package set

import "fmt"

func ArraysSum() {
	var arr1 = [5]int{7, 4, 13, 9, 2}
	var sum int

	for i := range arr1 {
		sum += arr1[i]
	}
	fmt.Println(arr1)
	fmt.Println(sum)
}
