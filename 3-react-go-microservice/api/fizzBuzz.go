package main

// FizzBuzzResult struct for holding the result of the FizzBuzz algorithm
type FizzBuzzResult struct {
	Input  int    `json:"input"`
	Output string `json:"output"`
}

func fizzBuzz(start int, end int) (result []FizzBuzzResult, err error) {

	result = make([]FizzBuzzResult, end)

	for x := start; x < end+1; x++ {

		var current *FizzBuzzResult = new(FizzBuzzResult)

		current.Input = x
		current.Output = ""

		if x%3 == 0 {
			current.Output += "Fizz"
		}
		if x%5 == 0 {
			current.Output += "Buzz"
		}

		result[x-1] = *current
	}

	return result, err
}
