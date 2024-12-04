package processing

import (
	"strconv"
	"strings"
)

func GetMeMyMuls(input string) []string {
	var listOfMuls []string
	var doOrDont bool = true // PART 2 - Checking if do & dont value
	for i := 0; i < len(input); i++ {
		if i+7 > len(input) {
			break //breaking if the i+7 is bigger than length of list 7 = number of chars in "don't()"
		}
		if string(input[i:i+7]) == "don't()" {
			doOrDont = false
		}
		if string(input[i:i+4]) == "do()" {
			doOrDont = true
		}
		if doOrDont == false {
			continue
		}
		if string(input[i:i+4]) == "mul(" {
			mulCheck := input[i : len(input)-1]                                         //Whole string from new start point
			if strings.Index(mulCheck, ")") >= 7 && strings.Index(mulCheck, ")") < 12 { //checking if closing bracket is at the right spot coz we know number of min and max digits
				mulCheck := input[i : i+strings.Index(mulCheck, ")")+1]
				if strings.Index(mulCheck, ",") > 1 {
					listOfMuls = append(listOfMuls, mulCheck) //list of formated strings
				}
			}
		}
	}
	return listOfMuls
}

func ProcessMyMuls(input []string) int {
	var result int
	for i := 0; i < len(input); i++ {
		if i > len(input) {
			break
		}
		indxpStart := strings.Index(input[i], "(")
		indxpEnd := strings.Index(input[i], ")")
		numbers := strings.Split(input[i][indxpStart+1:indxpEnd], ",") //extracting the numbers
		number1, _ := strconv.Atoi(numbers[0])                         //converting the numbers
		number2, _ := strconv.Atoi(numbers[1])                         //converting the numbers
		mulResult := number1 * number2
		result = mulResult + result
	}
	return result
}
