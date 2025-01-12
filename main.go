package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 1000
	screenHeight = 800
	blockSize    = 20 // Size of each block in the world
)

type Game struct {
	world [][]int
}

func generateFlatWorld(rows, cols int) [][]int {
	world := make([][]int, rows)
	for i := range world {
		world[i] = make([]int, cols)
		for j := range world[i] {
			world[i][j] = 1 // Fill world with blocks (1 = block present)
		}
	}
	return world
}

func NewGame() *Game {
	world := generateFlatWorld(10, 10) // Create a 10x10 world grid
	return &Game{world: world}
}

// Update is called every frame to update game state
func (g *Game) Update() error {
	return nil
}

// Draw is called every frame to render the game
func (g *Game) Draw(screen *ebiten.Image) {
	for y, row := range g.world {
		for x, block := range row {
			if block == 1 {
				blockX := float32(x * blockSize)
				blockY := float32(y * blockSize)
				// Draw a simple colored block with green color using vector.DrawFilledRect
				vector.DrawFilledRect(screen, blockX, blockY, float32(blockSize), float32(blockSize), color.RGBA{
					R: 76,  // Red component
					G: 179, // Green component
					B: 76,  // Blue component
					A: 255, // Alpha (opacity)
				}, false) // false for no anti-aliasing
			}
		}
	}
}

// Layout defines the screen layout (canvas size)
func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	// Initialize the game and start the game loop
	game := NewGame()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("GoCraft - Minecraft in Go")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
