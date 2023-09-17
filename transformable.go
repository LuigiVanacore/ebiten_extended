package ebiten_extended


type Transformable interface {
	GetTransform() *Transform
	SetTransform(transform Transform)
}