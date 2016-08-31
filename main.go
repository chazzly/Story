package main

// TODO: Sound
// TODO: Words for the story.
// TODO: Nav Bar  ( adjust click range to match)

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