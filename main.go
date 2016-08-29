package main

// TODO:  Reduce image sizes and scale them significantly less.
// TODO: Sound
// TODO: Words for the story.
// TODO: Nav Bar???
// TODO: Use clicks for forward and back.

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