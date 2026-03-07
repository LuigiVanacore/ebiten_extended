package main

import (
	"fmt"
	"image/color"
	"log"
	"path/filepath"

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/collision"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/tilemap"
	"github.com/LuigiVanacore/ebiten_extended/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 320
	screenHeight = 320
)

type Game struct {
	engine    *ebiten_extended.Engine
	tilemap   *tilemap.TileMapNode
	collMgr   *collision.CollisionManager
	pathfinder *tilemap.Pathfinder
	path      []math2D.Vector2D
	spawnTileX int
	spawnTileY int
	overlaps  []collision.CollisionParticipant
}

func NewGame() *Game {
	engine := ebiten_extended.NewEngine()
	engine.Input().SetMouseEnabled(true)
	engine.SetIsDebug(true) // show collision outlines on hover

	tmxPath := filepath.Join("example", "tilemap", "map.tmx")
	if _, err := tilemap.NewTileMapNode(tmxPath); err != nil {
		tmxPath = "map.tmx" // fallback: run from example/tilemap dir
	}
	tm, err := tilemap.NewTileMapNode(tmxPath)
	if err != nil {
		log.Fatalf("load tilemap: %v (run from example/tilemap with map.tmx and tiles.png)", err)
	}
	engine.World().AddNodeToDefaultLayer(tm)

	mask := collision.NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))
	if err := tm.BuildCollisionsFromObjectLayer("collisions", mask); err != nil {
		log.Printf("BuildCollisionsFromObjectLayer: %v", err)
	}
	collMgr := collision.NewCollisionManager()
	for _, col := range tm.GetChildren() {
		if c, ok := col.(*collision.Collider); ok {
			collMgr.AddParticipant(c)
		}
	}
	engine.World().SetPostUpdate(func() {
		collMgr.CheckCollision()
	})

	pf := tilemap.BuildPathfinderFromTileLayer(tm, "ground", true)
	if pf != nil {
		pf.SetAllowDiagonals(true)
	}

	tileW := float64(tm.GetMapData().TileWidth)
	tileH := float64(tm.GetMapData().TileHeight)
	spawnTileX, spawnTileY := 1, 1
	for _, og := range tm.GetMapData().ObjectGroups {
		if og.Name == "collisions" {
			for _, obj := range og.Objects {
				if obj.Name == "spawn" {
					spawnTileX = int(obj.X / tileW)
					spawnTileY = int(obj.Y / tileH)
					break
				}
			}
			break
		}
	}

	return &Game{
		engine:     engine,
		tilemap:    tm,
		collMgr:    collMgr,
		pathfinder: pf,
		spawnTileX: spawnTileX,
		spawnTileY: spawnTileY,
	}
}

func (g *Game) Update() error {
	cx, cy := g.engine.World().Camera().GetCursorCoords(g.engine.Input())
	g.overlaps = g.collMgr.OverlapPoint(math2D.NewVector2D(cx, cy))

	if g.pathfinder != nil && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		tileW := float64(g.tilemap.GetMapData().TileWidth)
		tileH := float64(g.tilemap.GetMapData().TileHeight)
		endX := int(cx / tileW)
		endY := int(cy / tileH)
		if endX >= 0 && endX < g.pathfinder.Width() && endY >= 0 && endY < g.pathfinder.Height() {
			nodes := g.pathfinder.FindPath(g.spawnTileX, g.spawnTileY, endX, endY)
			g.path = tilemap.PathToWorld(nodes, tileW, tileH)
		}
	}

	return g.engine.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.engine.Draw(screen)
	if g.engine.IsDebug() {
		for _, col := range g.overlaps {
			if c, ok := col.(*collision.Collider); ok {
				c.DrawDebug(screen, nil)
			}
		}
	}
	if len(g.path) >= 2 {
		cam := g.engine.World().Camera()
		for i := 0; i < len(g.path)-1; i++ {
			sx1, sy1 := cam.GetScreenCoords(g.path[i].X(), g.path[i].Y())
			sx2, sy2 := cam.GetScreenCoords(g.path[i+1].X(), g.path[i+1].Y())
			vector.StrokeLine(screen, float32(sx1), float32(sy1), float32(sx2), float32(sy2), 3, color.RGBA{255, 200, 0, 255}, true)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Tilemap - Collisions, OverlapPoint, A* Pathfinding (click to set path)")
	game := NewGame()
	fmt.Println("Tilemap loaded. Move mouse over collision objects. Set engine.SetIsDebug(true) to see outlines.")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
