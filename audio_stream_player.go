package ebiten_extended

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

// AudioStreamPlayer is a node that holds a reference to a cached sound and can Play/Pause/Stop it.
// Similar to Godot's AudioStreamPlayer: create once, keep a handle, trigger playback when needed.
// Implements SceneNode (via Node) and Updatable for loop support.
type AudioStreamPlayer struct {
	Node
	am      *AudioManager
	soundID string
	volume  float64
	loop    bool
	player  *audio.Player
}

// NewAudioStreamPlayer creates an AudioStreamPlayer for the given sound ID.
// The sound must already be loaded via AudioManager.AddSound.
// Returns an error if am is nil or the sound ID is not found.
// Typically use AudioManager.CreateStreamPlayer instead.
func NewAudioStreamPlayer(name, soundID string, am *AudioManager) (*AudioStreamPlayer, error) {
	if am == nil {
		return nil, errors.New("audio: AudioManager must not be nil")
	}
	if !am.HasSound(soundID) {
		return nil, errors.New("audio: sound not found: " + soundID)
	}
	return &AudioStreamPlayer{
		Node:    *NewNode(name),
		am:      am,
		soundID: soundID,
		volume:  1.0,
	}, nil
}

// Play starts or resumes playback. If already playing, has no effect.
func (a *AudioStreamPlayer) Play() {
	if a.am == nil || !a.am.HasSound(a.soundID) {
		return
	}
	if a.player != nil {
		if a.player.IsPlaying() {
			return
		}
		a.stopPlayer()
	}
	pcm, ok := a.am.getSoundPCM(a.soundID)
	if !ok {
		return
	}
	p := a.am.Context().NewPlayerF32FromBytes(pcm)
	p.SetVolume(a.volume * a.am.MasterVolume())
	p.Play()
	a.player = p
}

// Pause pauses playback if the player is active.
func (a *AudioStreamPlayer) Pause() {
	if a.player != nil && a.player.IsPlaying() {
		a.player.Pause()
	}
}

// Stop stops playback and releases the internal player. Call Play() to start again.
func (a *AudioStreamPlayer) Stop() {
	a.stopPlayer()
}

func (a *AudioStreamPlayer) stopPlayer() {
	if a.player == nil {
		return
	}
	_ = a.player.Close()
	a.player = nil
}

// IsPlaying reports whether the stream is currently playing.
func (a *AudioStreamPlayer) IsPlaying() bool {
	return a.player != nil && a.player.IsPlaying()
}

// SetVolume sets the volume in [0, 1]. Applied to the current player if playing.
func (a *AudioStreamPlayer) SetVolume(vol float64) {
	if vol < 0 {
		vol = 0
	}
	if vol > 1 {
		vol = 1
	}
	a.volume = vol
	if a.player != nil {
		a.player.SetVolume(vol * a.am.MasterVolume())
	}
}

// Volume returns the current volume.
func (a *AudioStreamPlayer) Volume() float64 {
	return a.volume
}

// SetLoop enables or disables looping. When true, playback restarts when it ends (handled in Update).
func (a *AudioStreamPlayer) SetLoop(loop bool) {
	a.loop = loop
}

// Loop returns whether looping is enabled.
func (a *AudioStreamPlayer) Loop() bool {
	return a.loop
}

// Update implements Updatable. Handles loop restart when playback finishes.
func (a *AudioStreamPlayer) Update() {
	if !a.loop || a.player == nil {
		return
	}
	if !a.player.IsPlaying() {
		a.stopPlayer()
		a.Play()
	}
}
