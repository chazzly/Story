package main

import (
	"log"
	"./structure"
	"github.com/hajimehoshi/ebiten"
)

func main() {
	game := structure.NewGame()
	update := game.Update
	if err := ebiten.Run(update, structure.ScreenWidth, structure.ScreenHeight, 2, "Story (Here we go!)"); err != nil {
		log.Fatal(err)
	}
}