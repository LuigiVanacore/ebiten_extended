package math2D
 

type Segment struct {
	startPoint Vector2D
	endPoint   Vector2D
}

func NewSegment(startPoint, endPoint Vector2D) Segment {
	return Segment{startPoint: startPoint, endPoint: endPoint}
}

func (segment *Segment) GetStartPoint() Vector2D {
	return segment.startPoint
}

func (segment *Segment) GetEndPoint() Vector2D {
	return segment.endPoint
}

func (segment *Segment) SetStartPoint(startPoint Vector2D) *Segment {
	segment.startPoint = startPoint
	return segment
}

func (segment *Segment) SetEndPoint(endPoint Vector2D) *Segment {
	segment.endPoint = endPoint
	return segment
}

func (s Segment) ProjectSegment(onto Vector2D, ontoIsUnit bool) Range {
	var ontoUnit Vector2D
	if ontoIsUnit {
		ontoUnit = onto
	} else {
		ontoUnit = onto.Normalize()
	}
	r := Range{}
	r.SetMinimun(DotProduct(ontoUnit, s.startPoint))
	r.SetMaximum(DotProduct(ontoUnit, s.endPoint))
	r = r.SortRange()
	return r
}
