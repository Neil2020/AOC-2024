package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type SimilarityType struct {
	number    int
	frequency int
}
type SimilarityArray []SimilarityType

func main() {
	//Prob1()
	Prob2()
}

func Prob1() {

	//FileData
	fileData, err := os.ReadFile("input.txt")
	//fileData, err := os.ReadFile("input2.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	var lnumber []int
	var rnumber []int
	for _, line := range strings.Split(string(fileData), "\n") {
		lindNumber, li := strconv.Atoi(strings.Split(line, " ")[0])
		rindNumber, ri := strconv.Atoi(strings.Split(line, " ")[3])

		if li != nil {
			fmt.Println("Error converting string to int")
			return
		}
		if ri != nil {
			fmt.Println("Error converting string to int")
			return
		}

		lnumber = append(lnumber, lindNumber)
		rnumber = append(rnumber, rindNumber)
	}

	//Part 1 - Sorted Array Lowest to Highest
	sort.Slice(lnumber, func(i, j int) bool {
		return lnumber[i] < lnumber[j]
	})

	sort.Slice(rnumber, func(i, j int) bool {
		return rnumber[i] < rnumber[j]
	})

	fmt.Println(lnumber)
	fmt.Println(rnumber)

	//Find the Distances For Each Pair and add to New Array
	var distances []int
	for i := 0; i < len(lnumber); i++ {
		if i == len(lnumber) {
			break
		}
		var d int = lnumber[i] - rnumber[i]
		if d < 0 {
			distances = append(distances, d*-1)
		} else {
			distances = append(distances, d)
		}
	}
	fmt.Println(distances)

	var distanceSum int = 0
	for i := 0; i < len(distances); i++ {
		distanceSum = distances[i] + distanceSum
	}
	fmt.Println("Total Distance: ", distanceSum)
}

func Prob2() {
	//FileData
	fileData, err := os.ReadFile("input2.txt")
	//fileData, err := os.ReadFile("input2.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	var lnumber []int
	var rnumber []int
	for _, line := range strings.Split(string(fileData), "\n") {
		lindNumber, li := strconv.Atoi(strings.Split(line, " ")[0])
		rindNumber, ri := strconv.Atoi(strings.Split(line, " ")[3])

		if li != nil {
			fmt.Println("Error converting string to int")
			return
		}
		if ri != nil {
			fmt.Println("Error converting string to int")
			return
		}

		lnumber = append(lnumber, lindNumber)
		rnumber = append(rnumber, rindNumber)
	}
	fmt.Println(lnumber)
	fmt.Println(rnumber)

	//Similarity Score
	var Similarity SimilarityArray
	for i := 0; i < len(lnumber); i++ {
		if i == len(lnumber) {
			break
		}
		var s SimilarityType
		s.number = lnumber[i]
		for j := 0; j < len(rnumber); j++ {
			if j == len(rnumber) {
				break
			}
			if s.number == rnumber[j] {
				s.frequency++
			}
		}
		Similarity = append(Similarity, s)
	}
	fmt.Println(Similarity)

	//Get the result
	var result int = 0
	for i := 0; i < len(Similarity); i++ {
		if i == len(Similarity) {
			break
		}
		result = (Similarity[i].number * Similarity[i].frequency) + result
	}
	fmt.Println("Result: ", result)

}
