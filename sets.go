package main

import "fmt"

type Set struct {
	x       []int
	element int
}

// addElement adds elements to a set of real integers.
func (s *Set) addElement() {
	s.x = append(s.x, s.element)
	fmt.Println(s.x)
}

// findIndexAndRemoveElement finds an element's index in the set and deletes it from the set.
func (s *Set) findIndexAndRemoveElement() []int {
	// find index.
	index := -1
	for i, value := range s.x {
		if value == s.element {
			index = i
			break
		}
	}
	// delete element at that index
	s.x = append(s.x[:index], s.x[index+1:]...)
	return s.x
}

// checkElement checks if a given element is a member of the set.
func (s *Set) checkElement() bool {
	for _, value := range s.x {
		if value == s.element {
			return true
		}
	}
	return false
}

func main() {
	setOne := Set{
		x:       []int{1, 2, 3, 4, 5},
		element: 10,
	}
	fmt.Println(setOne)
	// addElement is used.
	setOne.addElement()
	// checkElement is used.
	setTwo := Set{
		x:       setOne.x,
		element: 4,
	}
	fmt.Println(setTwo)
	setTwo.checkElement()
	//	// findIndexAndRemoveElement is used.
	setTwo.findIndexAndRemoveElement()
	fmt.Println(setTwo)
}
