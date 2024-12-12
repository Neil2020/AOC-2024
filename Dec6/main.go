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
	potentialLoop          []xAndYLocation
	guardMovementDirection int
	hasExited              bool
	hasLooped              bool
}
type xAndYLocation struct {
	x int
	y int
}

func main() {
	fileDate, _ := os.ReadFile("./input/input.txt")
	rows := strings.Split(string(fileDate), "\n")
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
	//fmt.Println("Len:", len(Positions.positionsMoved), "\n", Positions.positionsMoved)
	var counterOfObs int
	//Loop Through Positions that the guard has been and then update the same to have obsticals
	for i := 0; i < len(Positions.positionsMoved); i++ {
		fmt.Println("itteration ", i, "out of ", len(Positions.positionsMoved))
		if i == 0 {
			continue
		}
		var tempPositions patrollingMAP
		tempPositions.maxX = Positions.maxX
		tempPositions.maxY = Positions.maxY
		tempPositions.obsticals = append(Positions.obsticals, Positions.positionsMoved[i])
		tempPositions.guard = Positions.positionsMoved[0]
		tempPositions.positionsMoved = append(tempPositions.positionsMoved, Positions.positionsMoved[0])
		tempPositions.guardMovementDirection = 1
		tempPositions.hasLooped = false
		tempPositions.hasExited = false
		tempPositions = AddingObjectLoopCreated(tempPositions)
		if tempPositions.hasLooped == true {
			counterOfObs++
		}
	}
	fmt.Println(counterOfObs)
}

func AddingObjectLoopCreated(Positions patrollingMAP) patrollingMAP {
	whileLoop := 2
	for whileLoop <= 2 {
		if Positions.hasExited {
			return Positions
		}
		if Positions.hasLooped {
			return Positions
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
	return Positions
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
		Positions.potentialLoop = nil
	} else {
		//Check to see if Looped
		Positions.potentialLoop = append(Positions.potentialLoop, newPosition)
		//fmt.Println("Org Moved locations:", Positions.positionsMoved, "\n Potential Loop:", Positions.potentialLoop)
		//fmt.Println("Path Taken:", Positions.positionsMoved, "\nPotential Loop:", Positions.potentialLoop)
		if len(Positions.potentialLoop) > 1 {
			for z, _ := range Positions.positionsMoved {
				if len(Positions.potentialLoop) == 0 {
					break
				}
				if Positions.positionsMoved[z] == Positions.potentialLoop[0] {
					for q, _ := range Positions.potentialLoop {
						//fmt.Println(len(Positions.potentialLoop))
						if len(Positions.potentialLoop) > 10000 {
							Positions.hasLooped = true
							break
						}
						if len(Positions.potentialLoop) <= 0 {
							break
						}
						if z+q > len(Positions.positionsMoved)-1 {
							break
						}
						if Positions.potentialLoop[q] != Positions.positionsMoved[z+q] {
							break
						} else {
							if len(Positions.potentialLoop)+z >= len(Positions.positionsMoved) {
								break
							}
							if Positions.potentialLoop[0] == Positions.positionsMoved[z] && Positions.positionsMoved[z] == Positions.potentialLoop[len(Positions.potentialLoop)-1] {
								fmt.Println(Positions.positionsMoved)
								Positions.hasLooped = true
								break
							}
						}
					}
				}
			}
		}

	}
	Positions.guard = newPosition
	return Positions

}
