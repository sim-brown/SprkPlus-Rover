package direction

import (
	"time"

	"../Maps"
	"gobot.io/x/gobot/platforms/sphero/ollie"
)

const Ms float64 = 0.4347826
const Wait time.Duration = 3000
const Interval int = 5

var Start time.Time
/* DRIVER */
var ball *ollie.Driver

func SetBall(mainBall *ollie.Driver) {
	ball = mainBall
}

func getGradeDirection(direction string) uint16 {
	switch direction {
	case "NE":
		return 45
	case "NW":
		return 315
	case "SE":
		return 135
	case "SW":
		return 225
	case "W":
		return 270
	case "E":
		return 90
	case "N":
		return 0
	case "S":
		return 180
	default:
		return 0
	}
}

func setPosition(direction string) {
	switch direction {
	case "NE":
		maps.MoveNorthEast()
		break
	case "NW":
		maps.MoveNorthWest()
		break
	case "SE":
		maps.MoveSouthEast()
		break
	case "SW":
		maps.MoveSouthWest()
		break
	case "W":
		maps.MoveWest()
		break
	case "E":
		maps.MoveEast()
		break
	case "N":
		maps.MoveNorth()
		break
	case "S":
		maps.MoveSouth()
		break
	default:
		maps.MoveNorth()
		break
	}
}

/* Movimento ball nella direzione selezionata */
func MoveInDirection(direction string, speed uint8) {
	/* Colore RGB - verde */
	ball.SetRGB(0, 128, 0)

	/* Movimento del robot */
	time.Sleep(Wait * time.Millisecond)
	ball.Roll(0, getGradeDirection(direction))
	time.Sleep(Wait / 2 * time.Millisecond)
	ball.Roll(speed, getGradeDirection(direction))

	Start := time.Now()
	_ = Start
	maps.PrintMap()
}
