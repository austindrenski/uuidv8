package uuidv8

// uuidv8 returns a UUIDv8 using 60 bits of a timestamp and 62 bits of a hash.
//
// [5.8.  UUID Version 8][5.8.]
//
//	UUIDv8 provides a format for experimental or vendor-specific use
//	cases.  The only requirement is that the variant and version bits
//	MUST be set as defined in Sections 4.1 and 4.2.  UUIDv8's uniqueness
//	will be implementation specific and MUST NOT be assumed.
//
//	The only explicitly defined bits are those of the version and variant
//	fields, leaving 122 bits for implementation-specific UUIDs.  To be
//	clear, UUIDv8 is not a replacement for UUIDv4 (Section 5.4) where all
//	122 extra bits are filled with random data.
//
//	Some example situations in which UUIDv8 usage could occur:
//
//	* An implementation would like to embed extra information within the
//	  UUID other than what is defined in this document.
//
//	* An implementation has other application and/or language
//	  restrictions that inhibit the use of one of the current UUIDs.
//
//	Appendix B provides two illustrative examples of custom UUIDv8
//	algorithms to address two example scenarios.
//
//	 0                   1                   2                   3
//	 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//	|                           custom_a                            |
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//	|          custom_a             |  ver  |       custom_b        |
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//	|var|                       custom_c                            |
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//	|                           custom_c                            |
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//
//		Figure 12: UUIDv8 Field and Bit Layout
//
//	custom_a:
//		The first 48 bits of the layout that can be filled as an
//		implementation sees fit.  Occupies bits 0 through 47 (octets 0-5).
//
//	ver:
//		The 4-bit version field as defined by Section 4.2, set to 0b1000
//		(8).  Occupies bits 48 through 51 of octet 6.
//
//	custom_b:
//		12 more bits of the layout that can be filled as an implementation
//		sees fit.  Occupies bits 52 through 63 (octets 6-7).
//
//	var:
//		The 2-bit variant field as defined by Section 4.1, set to 0b10.
//		Occupies bits 64 and 65 of octet 8.
//
//	custom_c:
//		The final 62 bits of the layout immediately following the var
//		field to be filled as an implementation sees fit.  Occupies bits
//		66 through 127 (octets 8-15).
//
// [B.1.  Example of a UUIDv8 Value (Time-Based)][B.1.]
//
//	This example UUIDv8 test vector utilizes a well-known 64-bit Unix
//	Epoch timestamp with 10 ns precision, truncated to the least
//	significant, rightmost bits to fill the first 60 bits of custom_a and
//	custom_b, while setting the version bits between these two segments
//	to the version value of 8.
//
//	The variant bits are set; and the final segment, custom_c, is filled
//	with random data.
//
//	Timestamp is Tuesday, February 22, 2022 2:22:22.000000 PM GMT-05:00,
//	represented as 0x2489E9AD2EE2E00 or 164555774200000000 (10 ns-steps).
//
//	-------------------------------------------
//	field     bits value
//	-------------------------------------------
//	custom_a  48   0x2489E9AD2EE2
//	ver        4   0x8
//	custom_b  12   0xE00
//	var        2   0b10
//	custom_c  62   0b00, 0xEC932D5F69181C0
//	-------------------------------------------
//	total     128
//	-------------------------------------------
//	final: 2489E9AD-2EE2-8E00-8EC9-32D5F69181C0
//
//		Figure 27: UUIDv8 Example Time-Based Illustrative Example
//
// [5.8.]: https://www.rfc-editor.org/rfc/rfc9562#name-uuid-version-8
// [B.1.]: https://www.rfc-editor.org/rfc/rfc9562#name-example-of-a-uuidv8-value-t
func uuidv8(t int64, h uint64) [16]byte {
	var b [16]byte

	// t &= 0x0FFF_FFFF_FFFF_FFFF // mask to 60 bits
	b[0] = byte(t >> 52)
	b[1] = byte(t >> 44)
	b[2] = byte(t >> 36)
	b[3] = byte(t >> 28)
	b[4] = byte(t >> 20)
	b[5] = byte(t >> 12)
	b[6] = byte(t>>8)&0b_0000_1111 | 0b_1000<<4
	b[7] = byte(t)

	// h &= 0x3FFF_FFFF_FFFF_FFFF // mask to 62 bits
	b[8] = byte(h>>56)&0b_0011_1111 | 0b_10<<6
	b[9] = byte(h >> 48)
	b[10] = byte(h >> 40)
	b[11] = byte(h >> 32)
	b[12] = byte(h >> 24)
	b[13] = byte(h >> 16)
	b[14] = byte(h >> 8)
	b[15] = byte(h)

	return b
}
