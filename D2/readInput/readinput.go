package readinput

import (
	"os"
	"strconv"
	"strings"
)

func ReadFileContents(inputfileName string) [][]int {
	fileData, err := os.ReadFile(inputfileName)
	if err != nil {
		panic(err)
	}
	Data := strings.Split(string(fileData), "\r\n") //Windows Specific Rubish
	var colOfInts [][]int
	for i := 0; i < len(Data); i++ {
		row := strings.Split(Data[i], " ")
		var nRow []int
		for z := 0; z < len(row); z++ {
			number, err := strconv.Atoi(row[z])
			if err != nil {
				panic(err)
			}
			nRow = append(nRow, number)
		}
		colOfInts = append(colOfInts, nRow)
	}
	return colOfInts
}
