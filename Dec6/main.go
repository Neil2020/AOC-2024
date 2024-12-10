package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type patrollingMAP struct {
	maxX                   int
	maxY                   int
	obsticals              []xAndYLocation
	guard                  xAndYLocation
	positionsMoved         []xAndYLocation
	guardMovementDirection int
	hasExited              bool
}
type xAndYLocation struct {
	x int
	y int
}

func main() {
	fileDate, _ := os.ReadFile("./input/sample.txt")
	rows := strings.Split(string(fileDate), "\r\n")
	//Create an index of object + start position + max x and y axisis of the map
	var Positions patrollingMAP
	Positions.maxY = len(rows[0])
	Positions.maxX = len(rows)
	Positions.guardMovementDirection = 1 //1 = up, 2= right, 3=down, 4= left // because easy to do increments rather than switching between string values :(
	Positions.hasExited = false
	for i := 0; i < len(rows); i++ {
		for z := 0; z < len(rows[i]); z++ {
			if string(rows[i][z]) == "#" {
				var e xAndYLocation
				e.x = i
				e.y = z
				Positions.obsticals = append(Positions.obsticals, e)
			}
			if string(rows[i][z]) == "^" {
				var e xAndYLocation
				e.x = i
				e.y = z
				Positions.guard = e
				Positions.positionsMoved = append(Positions.positionsMoved, e)
			}
		}
	}
	whileLoop := 2
	for whileLoop <= 2 {
		if Positions.hasExited {
			break
		}
		switch Positions.guardMovementDirection {
		case 1:
			Positions = MoveGuard(Positions, 1)
			continue
		case 2:
			Positions = MoveGuard(Positions, 2)
			continue
		case 3:
			Positions = MoveGuard(Positions, 3)
			continue
		case 4:
			Positions = MoveGuard(Positions, 4)
			continue
		}
	}
	fmt.Println("Len:", len(Positions.positionsMoved), "\n", Positions.obsticals, "\n", Positions.positionsMoved)
}

func MoveGuard(Positions patrollingMAP, direction int) patrollingMAP {
	if Positions.guard.x > Positions.maxX || Positions.guard.y > Positions.maxY || Positions.guard.x < 0 || Positions.guard.y < 0 {
		Positions.hasExited = true
		return Positions
	}
	if direction == 1 {
		var newPosition xAndYLocation
		newPosition.x = Positions.guard.x - 1
		newPosition.y = Positions.guard.y

		//Check if the new Position is an obstical
		if slices.Contains(Positions.obsticals, newPosition) {
			Positions.guardMovementDirection = Positions.guardMovementDirection + 1
			return Positions
		} else {
			if !slices.Contains(Positions.positionsMoved, newPosition) {
				Positions.positionsMoved = append(Positions.positionsMoved, newPosition)
			}
			Positions.guard = newPosition
			return Positions
		}
	}
	if direction == 2 {
		var newPosition xAndYLocation
		newPosition.x = Positions.guard.x
		newPosition.y = Positions.guard.y + 1

		//Check if the new Position is an obstical
		if slices.Contains(Positions.obsticals, newPosition) {
			Positions.guardMovementDirection = Positions.guardMovementDirection + 1
			return Positions
		} else {
			if !slices.Contains(Positions.positionsMoved, newPosition) {
				Positions.positionsMoved = append(Positions.positionsMoved, newPosition)
			}
			Positions.guard = newPosition
			return Positions
		}
	}
	if direction == 3 {
		var newPosition xAndYLocation
		newPosition.x = Positions.guard.x + 1
		newPosition.y = Positions.guard.y

		//Check if the new Position is an obstical
		if slices.Contains(Positions.obsticals, newPosition) {
			Positions.guardMovementDirection = Positions.guardMovementDirection + 1
			return Positions
		} else {
			if !slices.Contains(Positions.positionsMoved, newPosition) {
				Positions.positionsMoved = append(Positions.positionsMoved, newPosition)
			}
			Positions.guard = newPosition
			return Positions
		}
	}
	if direction == 4 {
		var newPosition xAndYLocation
		newPosition.x = Positions.guard.x
		newPosition.y = Positions.guard.y - 1

		//Check if the new Position is an obstical
		if slices.Contains(Positions.obsticals, newPosition) {
			Positions.guardMovementDirection = 1
			return Positions
		} else {
			if !slices.Contains(Positions.positionsMoved, newPosition) {
				Positions.positionsMoved = append(Positions.positionsMoved, newPosition)
			}
			Positions.guard = newPosition
			return Positions
		}
	}
	return Positions
}
