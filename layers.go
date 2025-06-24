package ebiten_extended

import (
	"github.com/LuigiVanacore/ebiten_extended/utils"
	"github.com/hajimehoshi/ebiten/v2"
)



func DrawNode( node Drawable, target *ebiten.Image, op *ebiten.DrawImageOptions) {
	node.Draw(target, op)
}

type Layers struct {
	layers []utils.Stack[func()]
}

func NewLayers(layersNum int) *Layers {
	return &Layers{ layers: make([]utils.Stack[func()], layersNum)}
}


func (l *Layers) AddNodeToLayer(layedIndex int, node Drawable, target *ebiten.Image, op ebiten.DrawImageOptions) {
	f := func ()  {
		node.Draw(target, &op)
	}
	l.layers[layedIndex].Push(f)
}


func (l *Layers) DrawLayers() {
    for i := range l.layers {
        for !l.layers[i].IsEmpty() {
            f, _ := l.layers[i].Pop()
            f()
        }
    }
}