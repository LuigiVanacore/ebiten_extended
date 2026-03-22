package ludum

import "testing"

func TestAudioManagerRemoveSound(t *testing.T) {
	am := &AudioManager{
		sounds: map[string][]byte{
			"jump": {1, 2, 3},
		},
	}

	if removed := am.RemoveSound("jump"); !removed {
		t.Fatal("expected RemoveSound to return true for existing sound")
	}
	if am.HasSound("jump") {
		t.Fatal("expected sound to be removed from cache")
	}
	if removed := am.RemoveSound("jump"); removed {
		t.Fatal("expected RemoveSound to return false for missing sound")
	}
}

func TestAudioManagerClearSounds(t *testing.T) {
	am := &AudioManager{
		sounds: map[string][]byte{
			"jump":  {1},
			"shoot": {2},
		},
	}

	removed := am.ClearSounds()
	if removed != 2 {
		t.Fatalf("expected 2 removed sounds, got %d", removed)
	}
	if len(am.sounds) != 0 {
		t.Fatalf("expected empty sounds map after clear, got %d", len(am.sounds))
	}
	if am.HasSound("jump") || am.HasSound("shoot") {
		t.Fatal("expected all sounds to be removed")
	}
}
