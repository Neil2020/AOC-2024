package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileData, _ := os.ReadFile("./input/input.txt")
	rows := strings.Split(string(fileData), "\n")
	var results int

	//arr of Operators
	for _, s := range rows {
		resultNo, _ := strconv.Atoi(strings.Split(s, ":")[0])
		var valuesToCheck []int
		//create arr of Ints
		for i, z := range strings.Split(strings.Split(s, ":")[1], " ") {
			if i != 0 {
				q, _ := strconv.Atoi(string(z))
				valuesToCheck = append(valuesToCheck, q)
			}
		}
		if checkValues(resultNo, valuesToCheck) {
			results = resultNo + results
		}
	}
	fmt.Println("Results:", results)
}

func checkValues(expectedResult int, values []int) bool {
	for i := 0; i < len(values); i++ {
		if values[0] == expectedResult && len(values) == 1 {
			return true
		}
		if values[0] > expectedResult {
			return false
		}

		if len(values) == 2 {
			if values[0]*values[1] == expectedResult || values[0]+values[1] == expectedResult {
				return true
			}
		}
		if len(values)-1 < i+1 {
			if values[0] == expectedResult && len(values) == 1 {
				return true
			}
			if values[0] > expectedResult {
				return false
			}
			return false
		}
		//values to be multipled && added
		var valuesMul []int
		var valuesAdd []int
		var valuesCat []int
		valuesMul = append(valuesMul, values...)
		valuesAdd = append(valuesAdd, values...)
		valuesCat = append(valuesCat, values...)

		valuesAdd[i+1] = (valuesAdd[i] + valuesAdd[i+1])
		valuesAdd = valuesAdd[i+1:]

		valuesMul[i+1] = (valuesMul[i] * valuesMul[i+1])
		valuesMul = valuesMul[i+1:]

		s1 := strconv.Itoa(valuesCat[i])
		s2 := strconv.Itoa(valuesCat[i+1])
		s3, _ := strconv.Atoi(s1 + s2)
		valuesCat[i+1] = s3
		valuesCat = valuesCat[i+1:]

		if checkValues(expectedResult, valuesAdd) {
			return true
		}
		if checkValues(expectedResult, valuesMul) {
			return true
		}
		if checkValues(expectedResult, valuesCat) {
			return true
		}
	}
	return false
}
