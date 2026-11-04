package pointers

import "fmt"

type Number struct {
	Value int
}

// Methods

// Object doesn`t change
func (n Number) GetNumber() int {
	return n.Value
}

// Object changes
func (n *Number) GetPNumber() int {
	return n.Value
}

// Functions

func GetFNumber(n Number) int {
	return n.Value
}

func GetFPNumber(n *Number) int {
	return n.Value
}

func Use() {
	num := Number{666}

	fmt.Println(num.GetNumber())
	fmt.Println(num.GetPNumber())

	fmt.Println(GetFNumber(num))
	fmt.Println(GetFPNumber(&num))
}
