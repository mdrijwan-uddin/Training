package set

import "fmt"

func MapPrint() {
	a := map[string]string{"Name": "Marcus", "Surname": "Tan", "Born": "1997"}
	fmt.Printf("a\t%v\n", a)

	b := make(map[string]string)
	b["Day"] = "13th"
	b["Month"] = "November"
	b["Year"] = "2024"

	fmt.Printf("b\t%v\n", b)
}
