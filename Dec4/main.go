package main

import (
	"fmt"
	"os"
	"strings"
)

func checkIfElements(Arr string) bool {
	for z := 0; z < len(Arr); z++ {
		if z+4 > len(Arr) {
			break
		}
		if Arr[z:z+4] == "XMAS" || Arr[z:z+4] == "SAMX" {
			return true
		}
	}
	return false
}

func main() {
	fileData, err := os.ReadFile("./input/sample.txt")
	if err != nil {
		panic(err)
	}
	rows := strings.Split(string(fileData), "\r\n")
	var count int = 0
	for i := 0; i < len(rows); i++ {
		var horiArr string = string(rows[i])
		if len(rows)+i >= 4 {
			//Create Virtical String
			for v := 0; v < len(rows[i]); v++ {
				if string(rows[i][v]) == "X" || string(rows[i][v]) == "S" {
					var virtArr string = ""
					var diagArr string = ""
					for vA := 0; vA == 4; vA++ {
						if vA == 0 {
							virtArr = string(rows[i+vA][v])
							diagArr = string(rows[i+vA][v])
						} else {
							virtArr = virtArr + string(rows[i+vA][v])
							diagArr = diagArr + string(rows[i+vA][v+1]) // Does not Work
						}
					}
					//Check for Virtical String
					if checkIfElements(string(virtArr)) {
						count++
					}
					//Create Diagonal String
					if checkIfElements(string(diagArr)) {
						count++
					}
				}
			}

		}

		//Check For horizontal String
		if checkIfElements(horiArr) {
			count++
		}
	}
	fmt.Println(count)
}
