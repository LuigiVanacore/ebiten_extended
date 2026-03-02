package ebiten_extended



// Updatable suggests an entity requires periodic processing per frame. 
type Updatable interface {
	// Update progresses logical routines unique to the owning object instance.
	Update()
}