package main

import (
	"gocraft/craft"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// Initialize the game and start the agame loop
	game := craft.NewGame()
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("GoCraft - Minecraft in Go")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
