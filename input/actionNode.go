package input




 

type ActionNode interface {
	isActionActive(buffer InputBuffer) bool
	//isActionActive(buffer EventBuffer) (bool, ActionResult)
}

 


type RealTimeNode struct {

}

type RealtimeKeyLeaf struct {

}



type RealtimeMouseLeaf struct {
}

type OrNode struct {
	lhs ActionNode
	rhs ActionNode
}

func NewOrNode(lhs, rhs ActionNode) OrNode {
	return OrNode{
		lhs: lhs,
		rhs: rhs,
	}
}

func (o OrNode) isActionActive(buffer InputBuffer) bool {
	return o.lhs.isActionActive(buffer) || o.rhs.isActionActive(buffer)
}


type AndNode struct {
	lhs ActionNode
	rhs ActionNode
}

func NewAndNode(lhs, rhs ActionNode) AndNode {
	return AndNode{
		lhs: lhs,
		rhs: rhs,
	}
}

func (a AndNode) isActionActive(buffer InputBuffer) bool {
	return a.lhs.isActionActive(buffer) && a.rhs.isActionActive(buffer)
}

type NotNode struct {
	action ActionNode
}

func NewNotNode(action ActionNode) NotNode {
	return NotNode{
		action: action,
	}
}



func (n NotNode) isActionActive(buffer InputBuffer) bool {
	return !n.action.isActionActive(buffer)
}

 
