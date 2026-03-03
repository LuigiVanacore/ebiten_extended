package main

import (
	"log"
	"os"

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

type Game struct {
	engine       *ebiten_extended.Engine
	musicPlayer  *ebiten_extended.AudioStreamPlayer // node per musica (loop)
	sfxPlayer    *ebiten_extended.AudioStreamPlayer // node per effetti
}

func NewGame() *Game {
	engine := ebiten_extended.NewEngine()
	am := engine.Audio()

	// Carica suoni (sound.wav, sound.ogg, sound.mp3)
	for _, path := range []string{"sound.wav", "sound.ogg", "sound.mp3"} {
		data, err := os.ReadFile(path)
		if err == nil {
			var audioFmt ebiten_extended.AudioFormat
			switch path {
			case "sound.wav":
				audioFmt = ebiten_extended.AudioFormatWAV
			case "sound.ogg":
				audioFmt = ebiten_extended.AudioFormatOGG
			case "sound.mp3":
				audioFmt = ebiten_extended.AudioFormatMP3
			}
			if err := am.AddSound("sfx", data, audioFmt); err != nil {
				log.Printf("AddSound %s: %v", path, err)
			} else {
				log.Printf("Loaded %s", path)
				break
			}
		}
	}

	// Crea nodi audio nello stile Godot
	musicPlayer := am.CreateStreamPlayer("music", "sfx")
	if musicPlayer != nil {
		musicPlayer.SetLoop(true)
		musicPlayer.SetVolume(0.5)
		engine.World().AddNodeToDefaultLayer(musicPlayer) // necessario per Update (loop)
	}

	sfxPlayer := am.CreateStreamPlayer("sfx_node", "sfx")

	return &Game{
		engine:      engine,
		musicPlayer: musicPlayer,
		sfxPlayer:   sfxPlayer,
	}
}

func (g *Game) Update() error {
	// Spazio: play one-shot (vecchio stile) oppure via nodo
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if g.sfxPlayer != nil {
			g.sfxPlayer.Play()
		} else {
			g.engine.Audio().PlaySound("sfx")
		}
	}
	// M: toggle musica
	if inpututil.IsKeyJustPressed(ebiten.KeyM) && g.musicPlayer != nil {
		if g.musicPlayer.IsPlaying() {
			g.musicPlayer.Pause()
		} else {
			g.musicPlayer.Play()
		}
	}
	return g.engine.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.engine.Draw(screen)
	// Draw hint
	// (In a real game you'd use a text renderer)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Audio - Space: SFX | M: Music (add sound.wav/ogg/mp3)")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
