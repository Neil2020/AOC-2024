package main

import (
	"d3/processing"
	"fmt"
	"os"
)

func main() {
	fileData, _ := os.ReadFile("./input/input.txt")
	//files in processing folder
	bruh := processing.GetMeMyMuls(string(fileData))
	result := processing.ProcessMyMuls(bruh)
	fmt.Println("Result: ", result)
}
