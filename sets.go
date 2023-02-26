package main

import "fmt"

func addElement(x []int, a int) {
	x = append(x, a)
	fmt.Println(x)
}

func main() {
	setOne := []int{1, 2, 3, 4, 5}
	fmt.Println(setOne)
	addElement(setOne, 10)
}
