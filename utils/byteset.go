package utils

// ByteSet is a bit field used for storing flags.
type ByteSet uint64

// Set returns the set with the given flag enabled.
func (c ByteSet) Set(flag ByteSet) ByteSet {
	return c | flag
}

// Clear returns the set with the given flag disabled.
func (c ByteSet) Clear(flag ByteSet) ByteSet {
	return c &^ flag
}

// Has reports whether the provided flag is set.
func (c ByteSet) Has(flag ByteSet) bool {
	return c&flag != 0
}
