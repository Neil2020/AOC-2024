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
	hasLooped              bool
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
	fmt.Println("Len:", len(Positions.positionsMoved), "\n", Positions.positionsMoved)
}

func AddingObjectLoopCreated(Positions patrollingMAP) bool {
	whileLoop := 2
	for whileLoop <= 2 {
		if Positions.hasExited {
			return false
		}
		if Positions.hasLooped {
			return true
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
	return false
}

func MoveGuard(Positions patrollingMAP, direction int) patrollingMAP {
	if Positions.guard.x >= Positions.maxX-1 || Positions.guard.y >= Positions.maxY-1 || Positions.guard.x < 0 || Positions.guard.y < 0 {
		Positions.hasExited = true
		return Positions
	}
	if direction == 1 {
		var newPosition xAndYLocation
		newPosition.x = Positions.guard.x - 1
		newPosition.y = Positions.guard.y

		Positions = newPositionCompare(Positions, newPosition)
	}
	if direction == 2 {
		var newPosition xAndYLocation
		newPosition.x = Positions.guard.x
		newPosition.y = Positions.guard.y + 1

		Positions = newPositionCompare(Positions, newPosition)
	}
	if direction == 3 {
		var newPosition xAndYLocation
		newPosition.x = Positions.guard.x + 1
		newPosition.y = Positions.guard.y
		Positions = newPositionCompare(Positions, newPosition)
	}
	if direction == 4 {
		var newPosition xAndYLocation
		newPosition.x = Positions.guard.x
		newPosition.y = Positions.guard.y - 1

		Positions = newPositionCompare(Positions, newPosition)
	}
	return Positions
}

func newPositionCompare(Positions patrollingMAP, newPosition xAndYLocation) patrollingMAP {
	if newPosition.x > Positions.maxX-1 || newPosition.y > Positions.maxY-1 || newPosition.x < 0 || newPosition.y < 0 {
		Positions.hasExited = true
		return Positions
	}
	//Check if the new Position is an obstical
	if slices.Contains(Positions.obsticals, newPosition) && Positions.guardMovementDirection <= 3 {
		Positions.guardMovementDirection = Positions.guardMovementDirection + 1
		return Positions
	}
	if slices.Contains(Positions.obsticals, newPosition) && Positions.guardMovementDirection == 4 {
		Positions.guardMovementDirection = 1
		return Positions
	}
	if !slices.Contains(Positions.positionsMoved, newPosition) {
		Positions.positionsMoved = append(Positions.positionsMoved, newPosition)
	} else {
		//Check to see if Looped
		var listCheck []xAndYLocation
		for i, _ := range Positions.positionsMoved {
			if newPosition == Positions.positionsMoved[i] {
				listCheck = append(listCheck, Positions.positionsMoved[i:]...)
				listCheck = append(listCheck, newPosition)
			}
		}
		for i := range len(listCheck) / 2 {
			fmt.Println("Check", listCheck)
			if listCheck[:len(listCheck)/2][i] == listCheck[len(listCheck)/2:][i] {
				fmt.Println(listCheck[:len(listCheck)/2][i], listCheck[len(listCheck)/2:][i])
			}
		}
	}
	Positions.guard = newPosition
	return Positions

}
