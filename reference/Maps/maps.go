/**
	Maps  ROBOT
**/

package Maps

import (
	"fmt"
	"os"
	"os/exec"
)

/* Dimensione mappa  */
const dimensionMap = 20
const square = "="

/* Array mappa con attuale posizione robot */
var Maps [dimensionMap][dimensionMap]string
var currentPositionX int
var currentPositionY int

/* Coordinate ostacoli */
type CoordObstacle struct {
	North     string
	South     string
	East      string
	West      string
	NorthWest string
	SouthWest string
	NorthEast string
	SouthEast string
}

/*Inizializzazione mappa con misure*/
func InitMap() {

	for i := 0; i < dimensionMap; i++ {
		for j := 0; j < dimensionMap; j++ {
			Maps[i][j] = square
		}
	}

	setPositiontRobot(int(dimensionMap/2), int(dimensionMap/2))
}

/* Set posizione robot */
func setPositiontRobot(x int, y int) {

	currentPositionX = x
	currentPositionY = y

	checkMapLimit()
}

/* Check limite mappa */
func checkMapLimit() {

	if currentPositionX > dimensionMap-1 {

		for i := 1; i < dimensionMap; i++ {
			for j := 0; j < dimensionMap; j++ {

				if i == dimensionMap-1 {
					Maps[i-1][j] = Maps[i][j]
					Maps[dimensionMap-1][j] = square
				} else {
					Maps[i-1][j] = Maps[i][j]
				}
			}
		}

		currentPositionX--

	}

	if currentPositionY > dimensionMap-1 {

		for i := 0; i < dimensionMap; i++ {
			for j := 1; j < dimensionMap; j++ {

				if j == dimensionMap-1 {
					Maps[i][j-1] = Maps[i][j]
					Maps[i][dimensionMap-1] = square
				} else {
					Maps[i][j-1] = Maps[i][j]
				}
			}
		}

		currentPositionY--

	}

	if currentPositionX < 0 {

		for i := dimensionMap - 1; i > 0; i-- {
			for j := dimensionMap - 1; j >= 0; j-- {

				if i == 1 {
					Maps[i][j] = Maps[i-1][j]
					Maps[0][j] = square
				} else {
					Maps[i][j] = Maps[i-1][j]
				}
			}
		}

		currentPositionX++

	}

	if currentPositionY < 0 {

		for i := dimensionMap - 1; i >= 0; i-- {
			for j := dimensionMap - 1; j > 0; j-- {

				if i == 0 {
					Maps[i][j] = Maps[i][j-1]
					Maps[0][j-1] = square
				} else {
					Maps[i][j] = Maps[i][j-1]
				}
			}
		}

		currentPositionY++
	}
}

/* Set ostacolo */
func SetObstacle() {

	Maps[currentPositionX][currentPositionY] = "#"
}

/* Get ostacoli su posizione robot */
func LookAround() CoordObstacle {

	var tempCurrentPositionX = currentPositionX
	var tempCurrentPositionY = currentPositionY

	if currentPositionX <= 0 {

		tempCurrentPositionX++
	}

	if currentPositionY <= 0 {

		tempCurrentPositionY++
	}

	if currentPositionX >= dimensionMap-1 {

		tempCurrentPositionX--
	}

	if currentPositionY >= dimensionMap-1 {

		tempCurrentPositionY--
	}

	coord := CoordObstacle{
		North:     Maps[tempCurrentPositionX-1][tempCurrentPositionY],
		South:     Maps[tempCurrentPositionX+1][tempCurrentPositionY],
		East:      Maps[tempCurrentPositionX][tempCurrentPositionY+1],
		West:      Maps[tempCurrentPositionX][tempCurrentPositionY-1],
		NorthWest: Maps[tempCurrentPositionX-1][tempCurrentPositionY-1],
		SouthWest: Maps[tempCurrentPositionX+1][tempCurrentPositionY-1],
		NorthEast: Maps[tempCurrentPositionX-1][tempCurrentPositionY+1],
		SouthEast: Maps[tempCurrentPositionX+1][tempCurrentPositionY+1]}

	return coord
}

/* Set movimento NORD */
func MoveNorth() {

	setPositiontRobot(currentPositionX-1, currentPositionY)
}

/* Set movimento SUD */
func MoveSouth() {

	setPositiontRobot(currentPositionX+1, currentPositionY)
}

/* Set movimento EST */
func MoveEast() {

	setPositiontRobot(currentPositionX, currentPositionY+1)
}

/* Set movimento OVEST */
func MoveWest() {

	setPositiontRobot(currentPositionX, currentPositionY-1)
}

/* Set movimento NORD-EST */
func MoveNorthEast() {

	setPositiontRobot(currentPositionX-1, currentPositionY+1)
}

/* Set movimento SUD-EST */
func MoveSouthEast() {

	setPositiontRobot(currentPositionX+1, currentPositionY+1)
}

/* Set movimento NORD-OVEST */
func MoveNorthWest() {

	setPositiontRobot(currentPositionX-1, currentPositionY-1)
}

/* Set movimento SUD-OVEST */
func MoveSouthWest() {

	setPositiontRobot(currentPositionX+1, currentPositionY-1)
}

/* Stampa mappa */
func PrintMap() {

	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	for i := 0; i < dimensionMap; i++ {

		fmt.Printf("\n")

		for j := 0; j < dimensionMap; j++ {
			fmt.Printf("%s ", Maps[i][j])
		}
	}

	fmt.Printf("\n\n## Posizione Robot ##\n")
	fmt.Printf("Posizione X: %d \n", currentPositionX)
	fmt.Printf("Posizione Y: %d \n\n", currentPositionY)
	fmt.Printf("OSTACOLI: %s \n", LookAround())
}
