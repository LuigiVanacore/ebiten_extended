package transform


type Transformable interface {
	GetTransform() *Transform
	SetTransform(transform Transform)
}