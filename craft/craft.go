package craft

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 800
	screenHeight = 600
	blockSize    = 40
)

const (
	BlockTypeGrass = iota
	BlockTypeDirt
	BlockTypeWater
)

type Block struct {
	BlockType int
}

type Game struct {
	world     [][]Block
	playerX   float64
	playerY   float64
	moveSpeed float64
}

func generateFlatWorld(rows, cols int) [][]Block {
	world := make([][]Block, rows)
	for i := range world {
		world[i] = make([]Block, cols)
		for j := range world[i] {

			if i == 0 {
				world[i][j] = Block{BlockType: BlockTypeGrass}
			} else if j%2 == 0 {
				world[i][j] = Block{BlockType: BlockTypeDirt}
			} else {
				world[i][j] = Block{BlockType: BlockTypeWater}
			}
		}
	}
	return world
}

func NewGame() *Game {
	world := generateFlatWorld(80, 60)
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
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) { // Move up
		g.playerY -= g.moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) { // Move down
		g.playerY += g.moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) { // Move left
		g.playerX -= g.moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) { // Move right
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
			blockX := x*blockSize - int(g.playerX)%blockSize
			blockY := y*blockSize - int(g.playerY)%blockSize

			var blockColor color.RGBA

			switch block.BlockType {
			case BlockTypeDirt:
				blockColor = color.RGBA{139, 69, 19, 255}
			case BlockTypeGrass:
				blockColor = color.RGBA{0, 255, 0, 255}
			case BlockTypeWater:
				blockColor = color.RGBA{0, 0, 255, 255}
			}

			vector.DrawFilledRect(screen, float32(blockX), float32(blockY), float32(blockSize), float32(blockSize), blockColor, false)

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
