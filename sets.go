package main

import "fmt"

// addElement adds elements to a set of real integers.
func addElement(x []int, a int) {
	x = append(x, a)
	fmt.Println(x)
}

// checkElement checks if a given element is a member of the set.
func checkElement(x []int, element int) bool {
	for _, value := range x {
		if value == element {
			return true
		}
	}
	return false
}

func main() {
	setOne := []int{1, 2, 3, 4, 5}
	fmt.Println(setOne)
	// addElement is used.
	addElement(setOne, 10)
	// checkElement is used.
	y := checkElement(setOne, 4)
	fmt.Println(y, "is the response.")
	z := checkElement(setOne, 12)
	fmt.Println(z, "is a wrong response.")
}
