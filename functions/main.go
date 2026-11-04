package functions

// Common
func Add(a float32, b float32) float32 {
	return a + b
}

// Shorten
func AddSh(a, b float32) float32 {
	return a + b
}

// Function inside Function
func AddFn(Add func(float64, float64) float64) float64 {
	return Add(0, 0)
}

// CLosing
func AddCl() func(float64, float64) float64 {
	res := 0.0
	return func(a float64, b float64) float64 {
		res = a + b
		return res
	}
}
