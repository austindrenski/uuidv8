package uuidv8

import (
	"encoding/binary"
	"math/bits"
)

// UUIDv8 constructs a UUIDv8 using a 64-bit timestamp and the 58 most significant bits of a 64-bit hash.
//
//	byte:    0        1        2        3        4        5        6        7
//	        +--------+--------+--------+--------+--------+--------+--------+--------+
//	bits:   |TTTTTTTT|TTTTTTTT|TTTTTTTT|TTTTTTTT|TTTTTTTT|TTTTTTTT|EEEETTTT|TTTTTTTT|
//	        +--------+--------+--------+--------+--------+--------+--------+--------+
//	bits:   |AATTTTHH|HHHHHHHH|HHHHHHHH|HHHHHHHH|HHHHHHHH|HHHHHHHH|HHHHHHHH|HHHHHHHH|
//	        +--------+--------+--------+--------+--------+--------+--------+--------+
//	byte:    8        9        10       11       12       13       14       15
//
//	T: timestamp (64 bits)
//	H: hash (58 bits)
//	E: version (4 bits)
//	A: variant (2 bits)
func UUIDv8(t uint64, h uint64) [16]byte {
	hi := 0b1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_0000_0000_0000_0000 & t
	hi |= 0b0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_1111_0000_0000_0000 & uint64(0b1000<<12)
	hi |= 0b0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_1111_1111_1111 & bits.RotateLeft64(t, 60)

	lo := 0b1100_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000 & uint64(0b10<<62)
	lo |= 0b0011_1100_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000 & bits.RotateLeft64(t, 58)
	lo |= 0b0000_0011_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111 & bits.RotateLeft64(h, 58)

	var u [16]byte
	binary.BigEndian.PutUint64(u[0:8], hi)
	binary.BigEndian.PutUint64(u[8:16], lo)
	return u
}

// Hash extracts the 58-bit hash encoded in a UUIDv8 constructed by uuidv8.UUIDv8.
//
//	byte:    0        1        2        3        4        5        6        7
//	        +--------+--------+--------+--------+--------+--------+--------+--------+
//	bits:   |TTTTTTTT|TTTTTTTT|TTTTTTTT|TTTTTTTT|TTTTTTTT|TTTTTTTT|EEEETTTT|TTTTTTTT|
//	        +--------+--------+--------+--------+--------+--------+--------+--------+
//	bits:   |AATTTTHH|HHHHHHHH|HHHHHHHH|HHHHHHHH|HHHHHHHH|HHHHHHHH|HHHHHHHH|HHHHHHHH|
//	        +--------+--------+--------+--------+--------+--------+--------+--------+
//	byte:    8        9        10       11       12       13       14       15
//
//	T: timestamp (64 bits)
//	H: hash (58 bits)
//	E: version (4 bits)
//	A: variant (2 bits)
func Hash(u [16]byte) uint64 {
	lo := binary.BigEndian.Uint64(u[8:16])

	h := 0b1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1100_0000 & bits.RotateLeft64(lo, 6)

	return h
}

// Timestamp extracts the 64-bit timestamp encoded in a UUIDv8 constructed by uuidv8.UUIDv8.
//
//	byte:    0        1        2        3        4        5        6        7
//	        +--------+--------+--------+--------+--------+--------+--------+--------+
//	bits:   |TTTTTTTT|TTTTTTTT|TTTTTTTT|TTTTTTTT|TTTTTTTT|TTTTTTTT|EEEETTTT|TTTTTTTT|
//	        +--------+--------+--------+--------+--------+--------+--------+--------+
//	bits:   |AATTTTHH|HHHHHHHH|HHHHHHHH|HHHHHHHH|HHHHHHHH|HHHHHHHH|HHHHHHHH|HHHHHHHH|
//	        +--------+--------+--------+--------+--------+--------+--------+--------+
//	byte:    8        9        10       11       12       13       14       15
//
//	T: timestamp (64 bits)
//	H: hash (58 bits)
//	E: version (4 bits)
//	A: variant (2 bits)
func Timestamp(u [16]byte) uint64 {
	hi := binary.BigEndian.Uint64(u[0:8])
	lo := binary.BigEndian.Uint64(u[8:16])

	t := 0b1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_0000_0000_0000_0000 & hi
	t |= 0b0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_1111_1111_1111_0000 & bits.RotateLeft64(hi, 4)
	t |= 0b0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_1111 & bits.RotateLeft64(lo, 6)

	return t
}
