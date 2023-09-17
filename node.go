package ebiten_extended

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Node struct {
	entity any
}

func NewNode(entity any) *Node {
	return &Node{entity: entity}
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