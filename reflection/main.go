package reflection

import (
	"fmt"
	"reflect"
)

type Human struct {
	name string
	age  int
}

func UseReflection() {
	h := Human{"Nikita", 24}
	fmt.Println(h)

	htype := reflect.TypeOf(h)
	fmt.Println(htype)        // reflection.Human
	fmt.Println(htype.Name()) // Human

	fmt.Println(reflect.ValueOf(h)) // {Nikita 24}

	fmt.Println(reflect.ValueOf(h).NumField()) // 2

	fmt.Println(reflect.ValueOf(h).Field(0)) // Nikita

	fmt.Println(reflect.ValueOf(h).Kind()) // struct
}
