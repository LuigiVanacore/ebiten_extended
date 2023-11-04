package ebiten_extended

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Node struct {
	id uint64
	name string
	entity any
}

func NewNode(entity any, name string) *Node {
	id := NodeManager().GetNextIdVal()
	return &Node{entity: entity, id: id, name: name}
}

func (n *Node) Update(dt float64) {
	if entity, ok := n.entity.(Updatable); ok {
		entity.Update(dt)
	} else {
		fmt.Println("entity is not updatable")
	}
}

func (n *Node) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	if entity, ok := n.entity.(Drawable); ok {
		entity.Draw(target, op)
	}
}