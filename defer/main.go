package deferr

import (
	"fmt"
)

func GetDefer() {
	// Biba and Boba
	defer fmt.Print(" and Boba")
	fmt.Print("Biba")
}
