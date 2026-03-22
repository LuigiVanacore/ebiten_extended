package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Rect struct {
	X, Y, W, H float64
}

func (r *Rect) Collides(o *Rect) bool {
	return r.X < o.X+o.W && r.X+r.W > o.X && r.Y < o.Y+o.H && r.Y+r.H > o.Y
}

type Game struct {
	Player   Rect
	Obstacle Rect
	Speed    float64
}

func (g *Game) Update() error {
	prevX, prevY := g.Player.X, g.Player.Y

	// Horizontal movement
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.Player.X -= g.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.Player.X += g.Speed
	}
	// Check horizontal collision
	if g.Player.Collides(&g.Obstacle) {
		g.Player.X = prevX
	}

	// Vertical movement
	prevY = g.Player.Y // Save after possible horizontal correction
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.Player.Y -= g.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.Player.Y += g.Speed
	}
	// Check vertical collision
	if g.Player.Collides(&g.Obstacle) {
		g.Player.Y = prevY
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw obstacle
	ebitenutil.DrawRect(screen, g.Obstacle.X, g.Obstacle.Y, g.Obstacle.W, g.Obstacle.H, color.RGBA{200, 0, 0, 255})
	// Draw player
	ebitenutil.DrawRect(screen, g.Player.X, g.Player.Y, g.Player.W, g.Player.H, color.RGBA{0, 200, 0, 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func main() {
	game := &Game{
		Player:   Rect{X: 50, Y: 50, W: 32, H: 32},
		Obstacle: Rect{X: 200, Y: 150, W: 100, H: 100},
		Speed:    4,
	}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Ebiten Rectangle Collision Example")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
