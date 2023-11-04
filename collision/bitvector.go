package collision

import "errors"

type Bitvector []byte

func IsBitSet(bitvector []byte, n int) (bool, error) { // returns true if bit n is set; false otherwise.
	if n > len(bitvector) {
		return false, errors.New("n is greater than the size of bitvector")
	}
	return bitvector[n/8]&(1<<(n%8)) > 0, nil
}

func Bitset(bitvector []byte, n int) error {
	if n > len(bitvector) {
		return errors.New("n is greater than the size of bitvector")
	}
	bitvector[n/8] |= (1 << (n % 8)) // |= is 'bitwise or equals', analogous to +=
	return nil
}

func Bitunset(bitvector []byte, n int) error {
	if n > len(bitvector) {
		return errors.New("n is greater than the size of bitvector")
	}
	bitvector[n/8] -= (1 << (n % 8))

	return nil
}
