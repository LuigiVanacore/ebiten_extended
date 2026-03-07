package ebiten_extended

// Updatable suggests an entity requires periodic processing per frame.
// Uses fixed 60 TPS (Ebiten default); timing logic should use FIXED_DELTA internally.
type Updatable interface {
	Update()
}