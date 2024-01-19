package ebiten_extended

// import "github.com/hajimehoshi/ebiten/v2"

// type Scene struct {
// 	name string
// 	scenes []*Scene
// 	nodes []*Node
// }

// func NewScene(name string) *Scene {
// 	return &Scene{ name: name , scenes: make([]*Scene, 0), nodes: make([]*Node, 0)}
// }

// func (s *Scene) AddNode(node *Node) {
// 	s.nodes = append(s.nodes, node)
// }

// func (s *Scene) AddScene(scene *Scene) {
// 	s.scenes = append(s.scenes, scene)
// }

// func (s *Scene) Update(dt float64) {
// 	for _, node := range s.nodes {
// 		node.Update(dt)
// 	}
// 	for _, scene := range s.scenes {
// 		scene.Update(dt)
// 	}
// }

// func (s *Scene) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
// 	for _, node := range s.nodes {
// 		node.Draw(target, op)
// 	}
// 	for _, scene := range s.scenes {
// 		scene.Draw(target, op)
// 	}
// }