package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/sphero/ollie"
	"gobot.io/x/gobot/platforms/sphero"

	"./reference/Prolog"

	"./reference/Direction"
	"./reference/Maps"
)

func main() {

	var isCollision = false
	var current_direction string

	//get parameters from console (SK-9885)
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := ollie.NewDriver(bleAdaptor)


	// Default configuration
	ball.EnableStopOnDisconnect()
	ball.SetBackLEDOutput(1)
	ball.SetStabilization(true)
	ball.SetRotationRate(1)

	/* collConf = sphero.CollisionConfig{
		Method: 0x01,
		Xt: 0x20,
		Yt: 0x20,
		Xs: 0x20,
		Ys: 0x20,
		Dead: 0x60}
 */
	ball.ConfigureCollisionDetection(sphero.CollisionConfig{
		Method: 0x00,
		Xt: 0x40,
		Yt: 0x40,
		Xs: 0x40,
		Ys: 0x40,
		Dead: 0x60})


	ball.On("collision", func(data interface{}) {
		isCollision = true

		/* Colore RGB - rosso */
		ball.SetRGB(255, 0, 0)

		/* Tempo di collisione */
		elapsed := time.Since(direction.Start)

		/* Cm percorsi */
		mRide := int((elapsed.Seconds() * direction.Ms) / 10)
		fmt.Printf("Tempo di collisione %d \n", mRide)

		for i := 0; i < mRide; i++ {
			direction.SetPosition(current_direction)
		}
		Maps.SetObstacle()
	})

	//setting ball to direction library
	direction.SetBall(ball)

	//map init
	Maps.InitMap()
	Maps.PrintMap()

	work := func() {
		for {
			direction.Start = time.Now()
			prolog.SetDirOfMap()

			speed, current_direction := prolog.MakeMove()

			if !isCollision {
				fmt.Println("is collision: ", isCollision)
				for i := 0; i < speed; i++ {
					direction.SetPosition(current_direction)
				}
			}
			isCollision = false

			prolog.Reset()
		}
	}

	if len(os.Args) > 2 {
		switch os.Args[2] {
		case "no-blt":
			for {
				work()
			}
		case "interactive":
			direction.Wait = 500
			work := func() {
				for {
					buf := bufio.NewReader(os.Stdin)
					fmt.Print("> ")
					sentence, _ := buf.ReadString('\n')
					sentence = strings.Replace(sentence, "\n", "", -1)
					fmt.Println(sentence)
					direction.MoveInDirection(string(sentence), 30)
				}
			}
			//New adapter
			robot := gobot.NewRobot("sprkplus",
				[]gobot.Connection{bleAdaptor},
				[]gobot.Device{ball},
				work)

			robot.Start()

			break
		}
	} else { //normal execution

		//New adapter
		robot := gobot.NewRobot("sprkplus",
			[]gobot.Connection{bleAdaptor},
			[]gobot.Device{ball},
			work,
		)

		robot.Start()
	}
}
