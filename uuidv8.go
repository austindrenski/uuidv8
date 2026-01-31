package uuidv8

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
func UUIDv8(t int64, h uint64) [16]byte {
	var u [16]byte

	u[0] = byte(t >> 56)
	u[1] = byte(t >> 48)
	u[2] = byte(t >> 40)
	u[3] = byte(t >> 32)
	u[4] = byte(t >> 24)
	u[5] = byte(t >> 16)
	u[6] = 0b1000_0000 | 0b0000_1111&byte(t>>12)
	u[7] = byte(t >> 4)
	u[8] = 0b1000_0000 | 0b0011_1100&byte(t<<2) | 0b0000_0011&byte(h>>62)
	u[9] = byte(h >> 54)
	u[10] = byte(h >> 46)
	u[11] = byte(h >> 38)
	u[12] = byte(h >> 30)
	u[13] = byte(h >> 22)
	u[14] = byte(h >> 14)
	u[15] = byte(h >> 6)

	return u
}
