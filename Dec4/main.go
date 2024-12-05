package main

import (
	"fmt"
	"os"
	"strings"
)

func checkIfElementsPart1(input string) bool {
	for z := 0; z < len(input); z++ {
		//Check min length
		if z+4 > len(input) {
			break
		}
		//Check to validate + return bool
		if input[z:z+4] == "XMAS" || input[z:z+4] == "SAMX" {
			return true
		}
	}
	return false
}
func Part1() {
	//Read Files
	fileData, err := os.ReadFile("./input/input.txt")
	if err != nil {
		panic(err)
	}
	rows := strings.Split(string(fileData), "\r\n")
	var count int = 0
	var ArrOfStrings []string
	for i := 0; i < len(rows); i++ {
		//Create Virtical String
		for v := 0; v < len(rows[i]); v++ {
			//Check to see if you need to create Arrays based on on the current value
			if string(rows[i][v]) == "X" || string(rows[i][v]) == "S" {
				//Add to Array Horizontal Line || no further check needed
				if len(rows[i][v:]) >= 4 {
					ArrOfStrings = append(ArrOfStrings, string(rows[i][v:v+4]))
				}
				//Virtical and Diagonal arrays need to have min 3 rows below
				if len(rows)-i > 3 {
					var virtArr string = ""
					var diagRArr string = ""
					var diagLArr string = ""
					//Loop to create the virtical + diagonal strings
					for vA := 0; vA < 4; vA++ {
						if vA == 0 {
							//initial string for 'X' OR 'S'
							virtArr = string(rows[i+vA][v])
							diagRArr = string(rows[i+vA][v])
							diagLArr = string(rows[i+vA][v])
						} else {
							//Append to virt and diagonal strings
							virtArr = virtArr + string(rows[i+vA][v])
							if len(rows[i+vA][v:]) > 3 {
								diagRArr = diagRArr + string(rows[i+vA][v+vA:][0])
							}
							if len(rows[i+vA][:v]) >= 3 {
								diagLArr = diagLArr + string(rows[i+vA][:v][len(rows[i+vA][:v-vA])])
							}
						}
					}
					//Add to Array for Virtical String
					ArrOfStrings = append(ArrOfStrings, string(virtArr))
					//Add to Array Diagonal Right String
					ArrOfStrings = append(ArrOfStrings, string(diagRArr))
					ArrOfStrings = append(ArrOfStrings, string(diagLArr))
				}
			}
		}
	}
	for _, value := range ArrOfStrings {
		//Checking Strings
		if checkIfElementsPart1(value) {
			count++
		}
	}
	fmt.Println("Count of XMAS: ", count)
}
func checkIfElementsPart2(inputOne string, inputTwo string) bool {
	//Check if the inputs provided creates a X-MAS one way or the other
	if (inputOne == "MAS" || inputOne == "SAM") && (inputTwo == "MAS" || inputTwo == "SAM") {
		return true
	}
	return false
}
func Part2() {
	//Read file
	fileData, err := os.ReadFile("./input/input.txt")
	if err != nil {
		panic(err)
	}
	//Get  rows of data
	rows := strings.Split(string(fileData), "\r\n")
	var count int = 0
	//Loop Through Rows
	for i := 0; i <= (len(rows) - 1); i++ {
		if (len(rows) - 1) <= i {
			break
		}
		//Skip first and last Row as cannot make an X-MAS with no rows above or below
		if i != 0 && i <= (len(rows)-1) {
			//Loop through value of Rows
			for z := 0; z < len(rows[i]); z++ {
				//if Value is "A" check the correct values the the row above and below and Check Elements to see if they make and X-MAS
				if string(rows[i][z]) == "A" && len(rows[i][:z]) >= 1 && len(rows[i][z:]) > 1 {
					if checkIfElementsPart2(string(string(rows[i-1][z-1])+string(rows[i][z])+string(rows[i+1][z+1])), string(string(rows[i+1][z-1])+string(rows[i][z])+string(rows[i-1][z+1]))) {
						count++
					}
				}
			}
		}
	}
	fmt.Println("Results: ", count)
}
func main() {
	Part1()
	Part2()
}
