package main

import (
	"image/color"
	"log"

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/collision"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/physics"
	"github.com/LuigiVanacore/ebiten_extended/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth   = 640
	screenHeight  = 480
	floorWidth    = 640 // collider must cover entire floor drawing
	floorHeight   = 40
)

var (
	playerColor   = color.RGBA{0, 200, 100, 255}
	obstacleColor = color.RGBA{200, 50, 50, 255}
	areaColor     = color.RGBA{100, 100, 255, 100}
	floorColor    = color.RGBA{80, 60, 40, 255}
)

type Game struct {
	engine         *ebiten_extended.Engine
	physicsWorld   *physics.PhysicsWorld
	collisionMgr   *collision.CollisionManager
	player         *physics.RigidBody2D
	obstacle       *physics.RigidBody2D
	floor          *physics.RigidBody2D
	area           *collision.Area2D
}

func NewGame() *Game {
	engine := ebiten_extended.NewEngine()
	physicsWorld := physics.NewPhysicsWorld()
	physicsWorld.Gravity = math2D.NewVector2D(0, 400)

	collisionMgr := collision.NewCollisionManager()
	mustRigidBody := func(shape collision.CollisionShape, mask collision.CollisionMask) *physics.RigidBody2D {
		rb, err := physics.NewRigidBody2D(shape, mask)
		if err != nil {
			panic(err)
		}
		return rb
	}
	mustArea := func(shape collision.CollisionShape, mask collision.CollisionMask) *collision.Area2D {
		area, err := collision.NewArea2D(shape, mask)
		if err != nil {
			panic(err)
		}
		return area
	}
	mustAddBody := func(body *physics.RigidBody2D) {
		if err := physicsWorld.AddRigidBody(body); err != nil {
			panic(err)
		}
	}

	mask := collision.NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))

	// Player (RigidBody with circle shape)
	playerShape := collision.NewCollisionCircle(math2D.NewCircle(math2D.ZeroVector2D(), 20))
	player := mustRigidBody(playerShape, mask)
	player.SetPosition(100, 50)
	player.UsesGravity = true
	player.AddChildren(ebiten_extended.NewDrawnCircle("player", math2D.ZeroVector2D(), 20, playerColor, true, 0))

	// Obstacle (RigidBody, static - zero velocity)
	obstacleShape := collision.NewCollisionRect(math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(80, 80)))
	obstacle := mustRigidBody(obstacleShape, mask)
	obstacle.SetPosition(320, 390)
	obstacle.UsesGravity = false
	obstacle.Static = true
	obstacle.AddChildren(ebiten_extended.NewDrawnRectangle("obstacle", math2D.ZeroVector2D(), math2D.NewVector2D(80, 80), obstacleColor, true, 0))

	// Floor (static - collider same size as drawing, perfectly overlapped)
	floorSize := math2D.NewVector2D(float64(floorWidth), float64(floorHeight))
	floorShape := collision.NewCollisionRect(math2D.NewRectangle(math2D.ZeroVector2D(), floorSize))
	floor := mustRigidBody(floorShape, mask)
	floor.SetPosition(float64(floorWidth)/2, float64(screenHeight)-float64(floorHeight)/2)
	floor.UsesGravity = false
	floor.Static = true
	floor.AddChildren(ebiten_extended.NewDrawnRectangle("floor", math2D.ZeroVector2D(), floorSize, floorColor, true, 0))

	// Area2D (sensor/trigger)
	areaShape := collision.NewCollisionRect(math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(150, 50)))
	area := mustArea(areaShape, mask)
	area.SetPosition(450, 380)
	area.AddChildren(ebiten_extended.NewDrawnRectangle("area", math2D.ZeroVector2D(), math2D.NewVector2D(150, 50), areaColor, true, 0))

	area.OnBodyEntered.Connect(nil, func(ev collision.Area2DBodyEvent) {
		// Trigger detected - e.g. collectible, damage zone
	})

	engine.World().AddNodeToLayer(player, 0)
	engine.World().AddNodeToLayer(obstacle, 0)
	engine.World().AddNodeToLayer(floor, 0)
	engine.World().AddNodeToLayer(area, 0)

	mustAddBody(player)
	mustAddBody(obstacle)
	mustAddBody(floor)

	collisionMgr.AddParticipant(player)
	collisionMgr.AddParticipant(obstacle)
	collisionMgr.AddParticipant(floor)
	collisionMgr.AddParticipant(area)

	engine.World().SetPostUpdate(func() {
		physicsWorld.Step(ebiten_extended.FIXED_DELTA)
		collisionMgr.CheckCollision()
	})

	return &Game{
		engine:       engine,
		physicsWorld: physicsWorld,
		collisionMgr: collisionMgr,
		player:       player,
		obstacle:     obstacle,
		floor:        floor,
		area:         area,
	}
}

func (g *Game) Update() error {
	// Input: jump / move (keyboard + gamepad)
	jump := ebiten.IsKeyPressed(ebiten.KeySpace)
	moveLeft := ebiten.IsKeyPressed(ebiten.KeyArrowLeft)
	moveRight := ebiten.IsKeyPressed(ebiten.KeyArrowRight)

	// Gamepad: first connected controller
	ids := ebiten.AppendGamepadIDs(nil)
	if len(ids) > 0 {
		id := ids[0]
		if ebiten.IsStandardGamepadLayoutAvailable(id) {
			jump = jump || ebiten.IsStandardGamepadButtonPressed(id, ebiten.StandardGamepadButtonRightBottom) // A / Cross
			lx := ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisLeftStickHorizontal)
			if lx < -0.3 {
				moveLeft = true
			}
			if lx > 0.3 {
				moveRight = true
			}
		} else {
			jump = jump || ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton0) // fallback: button 0
			lx := ebiten.GamepadAxisValue(id, 0)
			if lx < -0.3 {
				moveLeft = true
			}
			if lx > 0.3 {
				moveRight = true
			}
		}
	}

	if jump {
		v := g.player.GetVelocity()
		if v.Y() > -15 && v.Y() < 15 {
			g.player.ApplyImpulse(math2D.NewVector2D(0, -350))
		}
	}
	if moveLeft {
		g.player.ApplyImpulse(math2D.NewVector2D(-8, 0))
	}
	if moveRight {
		g.player.ApplyImpulse(math2D.NewVector2D(8, 0))
	}

	return g.engine.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.engine.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Physics - Keys: Arrows move, Space jump. Gamepad: left stick, A/Cross jump")
	ebiten.SetRunnableOnUnfocused(true) // Run even when window loses focus
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
