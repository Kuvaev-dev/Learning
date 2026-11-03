package values

import "fmt"

// Private
var a int = 666

// Public
var A int = 123

// Private
func declaredValue() (c int) {
	c = 888
	return
}

// Public
func GetValue() int {
	var b int = 777
	fmt.Println(a)
	fmt.Println(declaredValue())
	return b
}
