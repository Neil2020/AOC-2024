package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	BeforePgNo int
	AfterPgNo  int
}

// Check Sequences against rules
func CheckSequences(input []int) bool {
	fmt.Println(input)
	return true
}

func main() {
	fileData, err := os.ReadFile("./input/sample.txt")
	if err != nil {
		panic(err)
	}
	//Create Arr for Rules & Arr for Page Nos
	var Rules []Rule
	var PageSequence [][]int
	rows := strings.Split(string(fileData), "\r\n")
	//Loop to create Rules and page Sequences in int format
	for i := 0; i < len(rows); i++ {
		if strings.Index(rows[i], "|") > 1 {
			var rule Rule
			rulesNos := strings.Split(rows[i], "|")
			rule.BeforePgNo, _ = strconv.Atoi(rulesNos[0])
			rule.AfterPgNo, _ = strconv.Atoi(string(rulesNos[1]))
			Rules = append(Rules, rule)
		}
		if strings.Index(rows[i], ",") > 1 {
			strNos := strings.Split(rows[i], ",")
			var pageNos []int
			for z := 0; z < len(strNos); z++ {
				pg, _ := strconv.Atoi(strNos[z])
				pageNos = append(pageNos, pg)
			}
			PageSequence = append(PageSequence, pageNos)
		}
	}

	//Check Sequences for Rules
	for i := 0; i < len(PageSequence); i++ {
		CheckSequences(PageSequence[i])
	}
}
