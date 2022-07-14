package main

import (
	"time"
	"math/rand"
	"github.com/usrnmea/tic-tac-toe/game"
)

func main() {
	var (
		p1 game.Bot
		p2 game.RealPlayer
	)

	rand.Seed(time.Now().UnixNano())

	if rand.Intn(2) == 1 {
		game.NewGame(&p1, &p2)
	} else {
		game.NewGame(&p2, &p1)
	}
}
