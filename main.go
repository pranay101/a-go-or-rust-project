package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 800
	screenHeight = 600
	blockSize    = 40 // Size of each block in the world
)

type Game struct {
	world     [][]int
	playerX   float64
	playerY   float64
	moveSpeed float64
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
	return &Game{
		world:     world,
		playerX:   float64(screenWidth / 2),
		playerY:   float64(screenHeight / 2),
		moveSpeed: 5,
	}
}

// Update is called every frame to update game state
func (g *Game) Update() error {
	// Player movement
	if ebiten.IsKeyPressed(ebiten.KeyW) { // Move up
		g.playerY -= g.moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) { // Move down
		g.playerY += g.moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) { // Move left
		g.playerX -= g.moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) { // Move right
		g.playerX += g.moveSpeed
	}

	// Keep player within world bounds (optional)
	if g.playerX < 0 {
		g.playerX = 0
	}
	if g.playerY < 0 {
		g.playerY = 0
	}
	if g.playerX > float64(screenWidth-blockSize) {
		g.playerX = float64(screenWidth - blockSize)
	}
	if g.playerY > float64(screenHeight-blockSize) {
		g.playerY = float64(screenHeight - blockSize)
	}

	return nil
}

// Draw is called every frame to render the game
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the world with offset based on the player's position
	for y, row := range g.world {
		for x, block := range row {
			if block == 1 {
				blockX := x*blockSize - int(g.playerX)%blockSize
				blockY := y*blockSize - int(g.playerY)%blockSize
				// Draw the block (green color)
				blockColor := color.RGBA{0, 255, 0, 255} // Green block
				screen.Fill(blockColor)
				vector.DrawFilledRect(screen, float32(blockX), float32(blockY), float32(blockSize), float32(blockSize), blockColor, false)
			}
		}
	}

	// Draw the player (red color)
	playerColor := color.RGBA{255, 0, 0, 255} // Red player block
	vector.DrawFilledRect(screen, float32(g.playerX), float32(g.playerY), blockSize, blockSize, playerColor, false)
}

// Layout defines the screen layout (canvas size)
func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	// Initialize the game and start the agame loop
	game := NewGame()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("GoCraft - Minecraft in Go")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
