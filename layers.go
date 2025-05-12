package ebiten_extended

import "github.com/hajimehoshi/ebiten/v2"



func DrawNode( node Drawable, target *ebiten.Image, op *ebiten.DrawImageOptions) {
	node.Draw(target, op)
}

type Layers struct {
	layers [][]func()
}

func NewLayer(layersNum int) *Layers {
	return &Layers{ layers: make([][]func(), layersNum)}
}


func (l *Layers) AddNodeToLayer(layedIndex int, node Drawable, target *ebiten.Image, op *ebiten.DrawImageOptions) {
	f := func ()  {
		node.Draw(target, op)
	}
	l.layers[layedIndex] = append(l.layers[layedIndex], f)
}


func (l *Layers) DrawLayers() {
	for _, layer := range l.layers {
		for _, f := range layer {
			f()
		}
	}
}