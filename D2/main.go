package main

import (
	readinput "d2/readInput"
	"d2/rules"
	"fmt"
)

func main() {
	var ArrayOfInputs [][]int = readinput.ReadFileContents("./input/input.txt")
	var count int
	for _, iList := range ArrayOfInputs {
		if validateByMovingOne(iList) {
			count++
		}
	}
	fmt.Println("Solution - ", count)
}

func validate(input []int) bool {
	inc := rules.Increasing(input)
	dec := rules.Dncreasing(input)
	max := rules.MaxDiff(input, 3)
	min := rules.MinDiff(input, 1)
	return (inc || dec) && min && max
}

func validateByMovingOne(input []int) bool {
	for i := 0; i < len(input); i++ {
		data := append([]int{}, input[:i]...)
		data = append(data, input[i+1:]...)
		if validate(data) {
			return true
		}
	}
	return false
}
