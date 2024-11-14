package set

import "fmt"

func SliceAppend() {
	myslice2 := []string{"string1", "string2", "string3", "string4"}
	fmt.Printf("Before Append: ")
	fmt.Println(myslice2)

	myslice2 = append(myslice2, "string5")
	myslice2 = append(myslice2, "string6")

	fmt.Printf("After Append: ")
	fmt.Println(myslice2)
}
