package ebiten_extended

import (
	"bytes"
	"io"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

// DefaultSampleRate is the audio context sample rate (48000 Hz).
const DefaultSampleRate = 48000

// AudioFormat identifies the encoded audio format.
type AudioFormat string

const (
	AudioFormatWAV AudioFormat = "wav"
	AudioFormatOGG AudioFormat = "ogg"
	AudioFormatMP3 AudioFormat = "mp3"
)

// AudioManager handles loading and playback of sounds and music.
type AudioManager struct {
	ctx       *audio.Context
	sounds    map[string][]byte // id -> decoded F32 PCM bytes
	masterVol float64
}

// NewAudioManager creates an AudioManager. Uses the existing audio context if one was
// already created (e.g. by another Engine); otherwise creates one with DefaultSampleRate.
// Only one audio context can exist per process.
func NewAudioManager() *AudioManager {
	return NewAudioManagerWithSampleRate(DefaultSampleRate)
}

// NewAudioManagerWithSampleRate creates an AudioManager. Reuses the current audio context
// if it exists; otherwise creates a new one with the given sample rate.
func NewAudioManagerWithSampleRate(sampleRate int) *AudioManager {
	ctx := audio.CurrentContext()
	if ctx == nil {
		ctx = audio.NewContext(sampleRate)
	}
	return &AudioManager{
		ctx:       ctx,
		sounds:    make(map[string][]byte),
		masterVol: 1.0,
	}
}

// Context returns the underlying audio context (e.g. for IsReady).
func (a *AudioManager) Context() *audio.Context {
	return a.ctx
}

// AddSound decodes audio data and caches it for playback.
// Supports WAV, OGG (Vorbis), and MP3. Sample rate is resampled to match the context if needed.
func (a *AudioManager) AddSound(id string, data []byte, format AudioFormat) error {
	stream, err := a.decodeToStream(bytes.NewReader(data), format)
	if err != nil {
		return err
	}
	// Resample if necessary to match context sample rate
	var src io.Reader = stream
	size := stream.Length()
	from := stream.SampleRate()
	to := a.ctx.SampleRate()
	if from != to {
		src = audio.ResampleF32(stream, size, from, to)
	}
	pcm, err := io.ReadAll(src)
	if err != nil {
		return err
	}
	a.sounds[id] = pcm
	return nil
}

type stream interface {
	io.ReadSeeker
	Length() int64
	SampleRate() int
}

func (a *AudioManager) decodeToStream(r io.Reader, format AudioFormat) (stream, error) {
	switch format {
	case AudioFormatWAV:
		return wav.DecodeF32(r)
	case AudioFormatOGG:
		return vorbis.DecodeWithSampleRate(a.ctx.SampleRate(), r)
	case AudioFormatMP3:
		return mp3.DecodeF32(r)
	default:
		return nil, nil
	}
}

// PlaySound plays a cached sound once. Creates a new player each call.
// Volume is scaled by master volume.
func (a *AudioManager) PlaySound(id string) {
	a.PlaySoundWithVolume(id, 1.0)
}

// PlaySoundWithVolume plays a cached sound with volume in [0, 1].
func (a *AudioManager) PlaySoundWithVolume(id string, volume float64) {
	pcm, ok := a.sounds[id]
	if !ok {
		return
	}
	p := a.ctx.NewPlayerF32FromBytes(pcm)
	p.SetVolume(volume * a.masterVol)
	p.Play()
}

// SetMasterVolume sets the global volume multiplier [0, 1].
func (a *AudioManager) SetMasterVolume(vol float64) {
	if vol < 0 {
		vol = 0
	}
	if vol > 1 {
		vol = 1
	}
	a.masterVol = vol
}

// MasterVolume returns the current master volume.
func (a *AudioManager) MasterVolume() float64 {
	return a.masterVol
}

// HasSound reports whether the given sound ID is cached.
func (a *AudioManager) HasSound(id string) bool {
	_, ok := a.sounds[id]
	return ok
}

// CreateStreamPlayer creates an AudioStreamPlayer node for the given sound ID.
// The sound must already be loaded via AddSound. Returns nil if the sound is not found.
// Add the node to the World for loop support (Update), or use it standalone for one-shot control.
func (a *AudioManager) CreateStreamPlayer(name, soundID string) *AudioStreamPlayer {
	return NewAudioStreamPlayer(name, soundID, a)
}

// getSoundPCM returns the decoded PCM data for the given sound ID, or (nil, false) if not found.
// Used internally by AudioStreamPlayer.
func (a *AudioManager) getSoundPCM(id string) ([]byte, bool) {
	pcm, ok := a.sounds[id]
	return pcm, ok
}
