package math2d

// Range represents a 1D interval with minimum and maximum values.
type Range struct {
	minimum float64
	maximum float64
}

// NewRange returns a range with the given minimum and maximum.
func NewRange(minimum, maximum float64) Range {
	return Range{minimum: minimum, maximum: maximum}
}

func (r *Range) GetMinimum() float64 {
	return r.minimum
}

func (r *Range) GetMaximum() float64 {
	return r.maximum
}

func (r *Range) SetMinimum(minimum float64) *Range {
	r.minimum = minimum
	return r
}

func (r *Range) SetMaximum(maximum float64) *Range {
	r.maximum = maximum
	return r
}

func OverlappingRanges(a, b Range) bool {
	return b.minimum <= a.maximum && a.minimum <= b.maximum
}

func RangeHull(a, b Range) Range {
	var hull Range
	if a.minimum < b.minimum {
		hull.minimum = a.minimum
	} else {
		hull.minimum = b.minimum
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
	if r.minimum > r.maximum {
		sorted.minimum = r.maximum
		sorted.maximum = r.minimum
	}
	return sorted
}
