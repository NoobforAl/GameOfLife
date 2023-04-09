package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

type World [30][70]int8

func makeNewWorld30x30() World {
	return *new(World)
}

func (world *World) AddRandom() {
	for i := range *world {
		for j := range (*world)[i] {
			(*world)[i][j] = int8(rand.Intn(2))
		}
	}
}

func (world *World) Show() {
	symbol := []string{" ", "\u001b[32m*\u001b[0m"}
	fmt.Println("\033c")
	for x := 0; x < 30; x++ {
		for y := 0; y < 70; y++ {
			fmt.Print(symbol[(*world)[x][y]])
		}
		fmt.Println()
	}
}

func DeepEqual(emptyWorld, pastWorld *World) bool {
	if reflect.DeepEqual(emptyWorld, pastWorld) {
		return true
	}

	var i int
	for x := 0; x < 30; x++ {
		for y := 0; y < 70; y++ {
			if (*pastWorld)[x][y] != (*emptyWorld)[x][y] {
				i++
			}
		}
	}
	return i < 10
}

func (world *World) Action(pastWorld *World) {
	emptyWorld := makeNewWorld30x30()
	for x := 0; x < 30; x++ {
		for y := 0; y < 70; y++ {
			lives := 0
			for xd := x - 1; xd <= x+1; xd++ {
				for yd := y - 1; yd <= y+1; yd++ {
					if (*world)[(xd+30)%30][(yd+70)%70] == 1 {
						lives++
					}
				}
			}

			if (*world)[x][y] == 1 {
				lives--
			}

			if lives == 3 || (lives == 2 && (*world)[x][y] == 1) {
				emptyWorld[x][y] = 1
			} else {
				emptyWorld[x][y] = 0
			}
		}
	}

	if !DeepEqual(&emptyWorld, pastWorld) {
		for x := 0; x < 30; x++ {
			for y := 0; y < 70; y++ {
				(*world)[x][y] = emptyWorld[x][y]
				(*pastWorld)[x][y] = emptyWorld[x][y]
			}
		}
	} else {
		fmt.Println("Ops!")
		time.Sleep(3 * time.Second)
		fmt.Print("\033c")
		fmt.Println("We Need Destroy World!")
		fmt.Println("World Not Have Action!")
		fmt.Println("Restart World!")
		time.Sleep(3 * time.Second)
		world.AddRandom()
	}

}

func (world *World) Start() {
	past := makeNewWorld30x30()
	world.AddRandom()

	for {
		world.Action(&past)
		world.Show()
		time.Sleep(30 * time.Millisecond)
	}
}

func main() {
	life := makeNewWorld30x30()
	life.Start()
}
