package main

import (
	"errors"
	"image/color"
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	camera "github.com/LuigiVanacore/ebiten_extended"
)

// Organization is bad, but it's a messy example to test camera functions
var (
	cam    *camera.Camera
	tiles  *ebiten.Image
	player *ebiten.Image

	// When keyF is pressed, change follow mode (implemented by using
	// cam.SetPosition() or cam.MovePosition())
	CamFollowPlayer = true

	LastWindowWidth  int
	LastWindowHeight int

	LastMouseX      = 0
	LastMouseY      = 0
	MouseWasDown    bool
	MouseDownStartX int
	MouseDownStartY int
	MouseDownAt     time.Time
	MousePanAfter   = time.Millisecond * 100
	mx, my          float64 // mouse tile position
	px, py          int     // player tile position

	// Level and tile vars
	TileSize    = 100
	PlayerSize  = 75
	LevelWidth  = 30
	LevelHeight = 5
	Level       = []int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	}

	// Player vars
	PlayerX   float64 = 400.0
	PlayerY   float64
	PlayerRot float64
	VelX      float64
	VelY      float64
	Gravity   = 5.0
	JumpVel   = -40.0
	Jumping   = false

	ErrNormalExit = errors.New("Normal exit")
)

// Game required by ebiten
type Game struct{}

// Update updates the Game
func (g *Game) Update() error {
	PlayerRot += math.Pi / 200

	VelX = 0
	if !Jumping {
		VelY = 0
	}

	// Keyboard controls
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		CamFollowPlayer = !CamFollowPlayer
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyH) {
		VelX = -5
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyN) {
		VelX = 5
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if !Jumping {
			VelY = JumpVel
			Jumping = true
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ErrNormalExit
	}
	if ebiten.IsKeyPressed(ebiten.KeyG) {
		cam.Rotate(10)
	}
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		cam.Rotate(10)
	}

	// Physics
	if TileSize != 0 {
		VelY += Gravity

		// Cursor tile position
		mx, my = cam.GetCursorCoords()
		my = float64((int(my)) / int(TileSize))
		mx = float64((int(mx)) / int(TileSize))

		// Player tile position
		px = (int(PlayerX) + PlayerSize/2) / int(TileSize)
		py = (int(PlayerY) + PlayerSize) / int(TileSize)

		// Absolutely terrible collision detection and physics 🤫
		index := py*LevelWidth + px
		if index < LevelWidth*LevelHeight && index >= 0 {
			if Level[index] == 1 {
				// Touching tile
				Jumping = false
				VelY -= Gravity
				PlayerY = float64(py*TileSize) - float64(PlayerSize)
			}
		}

		PlayerX += VelX
		PlayerY += VelY

		if CamFollowPlayer {
			cam.SetPosition(PlayerX+float64(PlayerSize)/2, PlayerY+float64(PlayerSize)/2)
		}
	}

	// Panning, setting up for click events
	cx, cy := ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !MouseWasDown {
			// First frame mouse is down
			MouseWasDown = true
			MouseDownStartX = cx
			MouseDownStartY = cy
			MouseDownAt = time.Now()
		} else {
			// Pan when pressed for long enough
			if time.Now().Sub(MouseDownAt) > MousePanAfter && !CamFollowPlayer {
				cam.Translate(
					(float64(LastMouseX)-float64(cx))*1/cam.GetZoom(),
					(float64(LastMouseY)-float64(cy))*1/cam.GetZoom())
			}
		}

	} else if MouseWasDown {
		MouseWasDown = false
		// Only call mouse up event if the cursor didn't move more than a certain amount
		triggerMoveAmount := float64(TileSize) / 4.0
		if math.Abs(float64(MouseDownStartX-cx)) < triggerMoveAmount && math.Abs(float64(MouseDownStartY-cy)) < triggerMoveAmount {
			index := int(my)*LevelWidth + int(mx)
			if index >= 0 && index < LevelWidth*LevelHeight {
				switch Level[index] {
				case 0:
					Level[index] = 1
				case 1:
					Level[index] = 0
				}
			}
		}
	}

	LastMouseX = cx
	LastMouseY = cy

	// Zoom
	_, scrollAmount := ebiten.Wheel()
	if scrollAmount > 0 {
		cam.SetZoom(1.1)
	} else if scrollAmount < 0 {
		cam.SetZoom(0.9)
	}

	return nil
}

// Draw renders everything to screen
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw tiles image
	tiles = ebiten.NewImage(TileSize*LevelWidth, TileSize*LevelHeight)

	for y := 0; y < LevelHeight; y++ {
		for x := 0; x < LevelWidth; x++ {
			switch Level[y*LevelWidth+x] {
			case 0:
			case 1:
				ebitenutil.DrawRect(
					tiles,
					float64(x*TileSize),
					float64(y*TileSize),
					float64(TileSize),
					float64(TileSize),
					color.RGBA{0, 255, 0, 255})
			}
		}
	}
	if player == nil {
		// Draw player image
		player = ebiten.NewImage(PlayerSize, PlayerSize)
		player.Fill(color.RGBA{128, 0, 128, 255})
	}

	// Clear camera surface
	cam.Deallocate()
	cam.Fill(color.RGBA{255, 128, 128, 255})
	// Draw tiles
	tileOps := &ebiten.DrawImageOptions{}
	cam.DrawImage(tiles, cam.GetTranslation(tileOps, 0, 0))
	// Draw the player
	playerOps := &ebiten.DrawImageOptions{}
	playerOps = cam.GetTranslation(playerOps, PlayerX, PlayerY)
	cam.DrawImage(player, playerOps)

	// Draw to screen and zoom
	cam.Draw(screen)


}

// Layout sets window size
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	if LastWindowWidth != outsideWidth || LastWindowHeight != outsideHeight {
		cam.Resize(uint(outsideWidth), uint(outsideHeight))
		LastWindowWidth = outsideWidth
		LastWindowHeight = outsideHeight
	}
	return outsideWidth, outsideHeight
}

func main() {
	log.SetFlags(log.Lshortfile)

	w, h := 640*2, 480*2
	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle("Platformer example")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	cam = camera.NewCamera(uint(w), uint(h))

	game := &Game{}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}