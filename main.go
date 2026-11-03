package main

import (
	"fmt"
	"learning/valuesshortinit"
	// "matht"
	// "learning/export"
	// "learning/functions"
	// "learning/multireturn"
	// "learning/namedreturn"
	// "learning/valueswithinit"
)

// go mod init ...

func main() {
	// Hello World
	// fmt.Println("Hello World")

	// Format
	// fmt.Printf("The sum is: %f", 2.2+2.2)
	// fmt.Printf("The sum using Math is: %f", math.MaxInt64+math.MaxFloat64)

	// Export
	// fmt.Printf("The Exportable value is: %d\n", export.Exportable)
	// fmt.Printf("The Exportable Function is: %d\n", export.Export())

	// Simple Functions
	// fmt.Println(functions.Add(1, 2))
	// fmt.Println(functions.AddSh(3, 4))

	// Multireturn
	// biba, boba := multireturn.Return()
	// fmt.Println(biba, boba)

	// Named Return
	// fmt.Println(namedreturn.Return("Biba", "Boba"))

	// Values
	// fmt.Println(values.GetValue())
	// fmt.Println(values.A)

	// Values With Init
	// fmt.Println(valueswithinit.GetValues())

	// Values Short Init
	fmt.Println(valuesshortinit.GetValue())
}
