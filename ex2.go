package main

func Sum (inputs chan []int) chan []int {
	output := make(chan []int)
	// Create go routine to do the sum
	go func(x chan []int) {
		defer close(x)
		// Create results slice with size of inputs capacity for all Sum result.
		results := make([] int, cap(inputs))
		// Iterate over all input chanels
		for i := range inputs {
			// Iterate over values stored in input slice
			for j, v := range i {
				results[j] += v // Do the sum
			}
		}
		x <- results // Send results to output chan
	}(output)

	return output
}
