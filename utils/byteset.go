package utils


type ByteSet uint64

func (c ByteSet) Set(flag ByteSet) ByteSet {
	return c | flag
}

func (c ByteSet) Clear(flag ByteSet) ByteSet {
	return c &^ flag
}

func (c ByteSet) Has(flag ByteSet) bool {
	return c&flag != 0
}

