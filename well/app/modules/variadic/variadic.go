package variadic

func SimpleVariadicToSlic(numbers ...int) []int {
	return numbers
}

func MixedVariadicToSlic(name string, numbers ...int) (string, []int) {
	return name, numbers
}
