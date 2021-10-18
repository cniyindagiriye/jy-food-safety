package variadic_test

import (
	"testing"

	"github.com/cniyindagiriye/jy-food-safety/well/app/modules/variadic"
)

func TestSimpleVariadicToSlic(t *testing.T) {
	if val := variadic.SimpleVariadicToSlic(); val != nil {
		t.Error("value should be nil", nil)
	} else {
		t.Log("SimpleVariadicToSlic() -> nil")
	}

	// Test for random set of values
	vals := variadic.SimpleVariadicToSlic(1, 2, 3)
	expected := []int{1, 2, 3}
	isErr := false
	for i := 0; i < 3; i++ {
		if vals[i] != expected[i] {
			isErr = true
			break
		}
	}
	if isErr {
		t.Error("value should be []int{1,2,3}", vals)
	} else {
		t.Log("simpleVariadicToSlic(1, 2, 3) -> []int{1,2,3}")
	}

	vals = variadic.SimpleVariadicToSlic(expected...)
	isErr = false
	for i := 0; i < 3; i++ {
		if vals[i] != expected[i] {
			isErr = true
			break
		}
	}
	if isErr {
		t.Error("value should be []int{1, 2, 3}", vals)
	} else {
		t.Log("simpleVariadicToSlice([]int{1, 2, 3}...) ->	[]int{1, 2, 3}")
	}
}

func TestMixedVariadicToSlice(t *testing.T) {
	// Test for simple argument & no variadic arguments
	name, numbers := variadic.MixedVariadicToSlic("Bob")
	if name == "Bob" && numbers == nil {
		t.Log("Recieved as expected: Bob, <nil slice>")
	} else {
		t.Errorf("Received unexpected values: %v, %v", name, numbers)
	}
}
