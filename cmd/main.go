package main

import (
	"github.com/Edu4rdoNeves/space-war/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
