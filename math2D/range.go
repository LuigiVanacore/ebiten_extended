package math2D
 
 


type Range struct {
	minimun float64
	maximum float64
}

func NewRange(minimum, maximum float64) Range {
	return Range{ minimun: minimum, maximum: maximum}
}

func (r *Range) GetMinimun() float64 {
	return r.minimun
}

func (r *Range) GetMaximum() float64 {
	return r.maximum
}

func (r *Range) SetMinimun(minimun float64) *Range {
	r.minimun = minimun
	return r
}

func (r *Range) SetMaximum(maximum float64) *Range {
	r.maximum = maximum
	return r
}

func OverlappingRanges(a, b Range) bool {
	return b.minimun <= a.maximum && a.minimun <= b.maximum
}

func RangeHull(a, b Range) Range {
	var hull Range 
	if a.minimun < b.minimun {
		hull.minimun = a.minimun
	} else {
		hull.minimun = b.minimun
	}
	if a.maximum > b.maximum {
		hull.maximum = a.maximum
	} else {
		hull.maximum = b.maximum
	}
	return hull
}

func (r Range) SortRange() Range {
	sorted := r
	if r.minimun > r.maximum {
		sorted.minimun = r.maximum
		sorted.maximum = r.minimun
	}
	return sorted
}

