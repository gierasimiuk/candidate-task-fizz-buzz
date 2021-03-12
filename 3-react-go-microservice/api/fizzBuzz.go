package main

import "errors"

type FizzBuzzResult struct {
	Input int `json:"input"`
	Output string `json:"output"`
}

func fizzBuzz(start int, end int) (result []FizzBuzzResult, err error) {
	result = nil
	err = errors.New("TODO: Implement fizz buzz")

	return
}