package mymap

import "fmt"

type Number struct {
	num float64
}

func GetMap() {
	// Key - Value
	testMap := make(map[float64]Number)
	testMap[1] = Number{2}
	// fmt.Println(testMap[1])

	// Literals
	secondTestMap := map[string]Number{
		"First":  {4},
		"Second": {6},
	}
	// PrintMap(secondTestMap)

	// Operations
	Operations(secondTestMap)
}

func PrintMap(myMap map[string]Number) {
	for key, value := range myMap {
		fmt.Println("Key:", key, "Value:", value.num)
	}
}

func Operations(myMap map[string]Number) {
	// Add
	myMap["Third"] = Number{3}
	PrintMap(myMap)
	fmt.Println()
	// Update
	myMap["First"] = Number{1}
	PrintMap(myMap)
	fmt.Println()
	// Get
	fmt.Println(myMap["First"])
	fmt.Println()
	// Delete
	delete(myMap, "First")
	PrintMap(myMap)
	fmt.Println()
}
