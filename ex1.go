package main

import "reflect"

//Reverse reverses a slice.
func Reverse(slice interface{}) {
	var sliceElem reflect.Value
	// Check if slice is a pointer, when we receive *slice[]
	if reflect.TypeOf(slice).Kind() == reflect.Ptr {
		// First convert the Pointer into a Value, this gives &slice[]
		sliceElem = reflect.ValueOf(slice)
		// Then extract the actual value form the Value, this gives slice[]
		sliceElem = sliceElem.Elem()
	} else {
		// As we have received a concrete value (slice[])
		// transform it to Value as we want to call Interface() later.
		sliceElem = reflect.ValueOf(slice)
	}
	swap := reflect.Swapper(sliceElem.Interface())
	for i := 0; i < sliceElem.Len() / 2; i++ {
		// When i = 0, swap first index with last index of the slice
		// When i = 1, swap second index with second last index of the slice
		swap(i, (sliceElem.Len() - 1) - i)
	}
}
