package ebiten_extended

import (
	"fmt"

	"github.com/LuigiVanacore/ebiten_extended/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

// ErrInvalidLayerIndex is returned when a negative layer index is passed.
var ErrInvalidLayerIndex = fmt.Errorf("layer index cannot be negative")

func DrawNode(node Drawable, target *ebiten.Image, op *ebiten.DrawImageOptions) {
	node.Draw(target, op)
}

// Layers is a stack-based draw system. Layer indices define draw order (lower = drawn first).
type Layers struct {
	layers []utils.Stack[func()]
}

// NewLayers returns a new Layers instance.
func NewLayers() *Layers {
	return &Layers{layers: make([]utils.Stack[func()], 10)}
}

// AddNodeToLayerF adds a draw callback to the layer at the given index.
// Returns ErrInvalidLayerIndex if layerIndex is negative.
func (l *Layers) AddNodeToLayerF(layerIndex int, f func()) error {
	if layerIndex < 0 {
		return ErrInvalidLayerIndex
	}
	for layerIndex >= len(l.layers) {
		l.layers = append(l.layers, utils.Stack[func()]{})
	}
	l.layers[layerIndex].Push(f)
	return nil
}

// AddNodeToLayer adds a node's draw callback to the layer at the given index.
// Returns ErrInvalidLayerIndex if layerIndex is negative.
func (l *Layers) AddNodeToLayer(layerIndex int, node Drawable, target *ebiten.Image, op ebiten.DrawImageOptions) error {
	if layerIndex < 0 {
		return ErrInvalidLayerIndex
	}
	for layerIndex >= len(l.layers) {
		l.layers = append(l.layers, utils.Stack[func()]{})
	}
	f := func() {
		node.Draw(target, &op)
	}
	l.layers[layerIndex].Push(f)
	return nil
}



// DrawLayers executes all queued draw callbacks in layer order.
func (l *Layers) DrawLayers() {
    for i := range l.layers {
        for !l.layers[i].IsEmpty() {
            f, _ := l.layers[i].Pop()
            f()
        }
    }
}